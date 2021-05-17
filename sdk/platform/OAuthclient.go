package platform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gofynd/fdk-client-golang/sdk/common"
	"github.com/gofynd/fdk-client-golang/sdk/common/utils"
)

//FdkTokenIssueError ...
type FdkTokenIssueError struct {
	Message string `json:"message"`
}

//FdkOAuthCodeError ...
type FdkOAuthCodeError struct {
	Message string `json:"message"`
}

// NewFdkTokenIssueError constructs and returns new FdkTokenIssueError object
func NewFdkTokenIssueError(message string) *FdkTokenIssueError {
	return &FdkTokenIssueError{
		Message: message,
	}
}

// Error returns technical error message
func (f *FdkTokenIssueError) Error() string {
	return fmt.Sprintf("%s", f.Message)
}

// NewFdkOAuthCodeError constructs and returns new FdkOAuthCodeError object
func NewFdkOAuthCodeError(message string) *FdkOAuthCodeError {
	return &FdkOAuthCodeError{
		Message: message,
	}
}

// Error returns technical error message
func (f *FdkOAuthCodeError) Error() string {
	return fmt.Sprintf("%s", f.Message)
}

//RawToken holds raw token details
type RawToken struct {
	ExpiresIn    int                    `json:"expires_in"`
	AccessToken  string                 `json:"access_token"`
	RefreshToken string                 `json:"refresh_token"`
	TokenType    string                 `json:"token_type"`
	CurrentUser  map[string]interface{} `json:"current_user"`
}

//OAuthClient holds OAuth Client details
type OAuthClient struct {
	RawToken       RawToken
	Config         *PlatformConfig
	Token          string
	RefreshToken   string
	TokenExpiresIn int
}

//Option holds the authorization options
type Option struct {
	Scope       []string `json:"scope"`
	RedirectURI string   `json:"redirectUri"`
	State       string   `json:"state"`
	AccessMode  string   `json:"access_mode"`
}

//NewOAuthClient return OAuthClient instance
func NewOAuthClient(config *PlatformConfig) *OAuthClient {
	return &OAuthClient{
		Config: config,
	}
}

//GetAccessToken returns the access token
func (o *OAuthClient) GetAccessToken() string {
	return o.Token
}

//SetAccessToken sets access and refresh token
func (o *OAuthClient) SetAccessToken(rawtoken RawToken) {
	o.RawToken = rawtoken
	o.TokenExpiresIn = rawtoken.ExpiresIn
	o.Token = rawtoken.AccessToken
	if rawtoken.RefreshToken != "" {
		o.RefreshToken = rawtoken.RefreshToken
		o.retryOAuthToken(rawtoken.ExpiresIn)
	}
	return
}

func (o *OAuthClient) retryOAuthToken(expiresIn int) {
	if expiresIn > 60 {
		go func() {
			select {
			case <-time.After(time.Duration(expiresIn-60) * time.Second):
				o.renewAccessToken()
			}
		}()
	}
}

//StartAuthorization perform user authorization with the provided scope
func (o *OAuthClient) StartAuthorization(options Option) (string, error) {
	var (
		queryParams = make(map[string]string)
		payload     = &strings.Reader{}
		params      = url.Values{}
	)

	queryParams["client_id"] = o.Config.APIKey
	queryParams["scope"] = strings.Join(options.Scope, ",")
	queryParams["redirect_uri"] = options.RedirectURI
	queryParams["state"] = options.State
	queryParams["access_mode"] = options.AccessMode
	queryParams["response_type"] = "code"

	u, err := url.Parse(o.Config.Domain)
	if err != nil {
		return "", NewFdkOAuthCodeError(err.Error())
	}
	apiURL := fmt.Sprintf("/service/panel/authentication/v1.0/company/%s/oauth/authorize", o.Config.CompanyID)

	//Parsing API URL
	u.Path = path.Join(u.Path, apiURL)
	apiURL = u.String()

	//Adding query params
	for k, v := range queryParams {
		params.Add(k, v)
	}

	//Converting query params to decoded query string
	decodedQueryString, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", NewFdkOAuthCodeError(err.Error())
	}
	//Creating HTTP Request
	req, err := http.NewRequest("GET", apiURL, payload)
	if err != nil {
		return "", NewFdkOAuthCodeError(err.Error())
	}

	//Setting query params to http request URL
	req.URL.RawQuery = decodedQueryString

	//Logging curl command
	// command, _ := http2curl.GetCurlCommand(req)
	// log.Println("curl = ", command)

	reqSigner := common.NewRequestSigner(req, true)
	err = reqSigner.Sign()
	if err != nil {
		return "", NewFdkOAuthCodeError(err.Error())
	}
	decodedQueryString, err = url.QueryUnescape(reqSigner.UpdatedReq.URL.Query().Encode())
	return fmt.Sprintf("%s/service/panel/authentication/v1.0/company/%s/oauth/authorize?%s", o.Config.Domain, o.Config.CompanyID, decodedQueryString), nil
}

//Query holds the authorization code generated after user grant permission
type Query struct {
	Code string `json:"code"`
}

//VerifyCallback performs API call to get access and refresh token
func (o *OAuthClient) VerifyCallback(query Query) error {
	if query.Code == "" {
		return NewFdkOAuthCodeError("Invalid authorization code")
	}
	token := utils.EncodeToBase64(fmt.Sprintf("%s:%s", o.Config.APIKey, o.Config.APISecret))
	apiURL := fmt.Sprintf("/service/panel/authentication/v1.0/company/%s/oauth/token", o.Config.CompanyID)

	headers := make(map[string]string)
	body := make(map[string]interface{})

	//Setting headers
	headers["Authorization"] = "Basic " + token
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	//Setting request body
	body["grant_type"] = "authorization_code"
	body["code"] = query.Code

	rawReq := &RawRequest{
		BaseURL:     o.Config.Domain,
		APIURL:      apiURL,
		Method:      "POST",
		Headers:     headers,
		QueryParams: nil,
		Body:        body,
		EncodeData:  true,
	}
	byteResponse, err := rawReq.Execute()
	if err != nil {
		return NewFdkTokenIssueError(err.Error())
	}
	rawToken := RawToken{CurrentUser: make(map[string]interface{})}
	err = json.Unmarshal(byteResponse, &rawToken)
	if err != nil {
		return NewFdkTokenIssueError(err.Error())
	}
	o.SetAccessToken(rawToken)
	return nil
}

func (o *OAuthClient) renewAccessToken() error {

	token := utils.EncodeToBase64(fmt.Sprintf("%s:%s", o.Config.APIKey, o.Config.APISecret))
	apiURL := fmt.Sprintf("/service/panel/authentication/v1.0/company/%s/oauth/token", o.Config.CompanyID)

	headers := make(map[string]string)
	body := make(map[string]interface{})

	//Setting headers
	headers["Authorization"] = "Basic " + token
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	//Setting request body
	body["grant_type"] = "refresh_token"
	body["refresh_token"] = o.RefreshToken

	rawReq := &RawRequest{
		BaseURL:     o.Config.Domain,
		APIURL:      apiURL,
		Method:      "POST",
		Headers:     headers,
		QueryParams: nil,
		Body:        body,
		EncodeData:  true,
	}
	byteResponse, err := rawReq.Execute()
	if err != nil {
		return NewFdkTokenIssueError(err.Error())
	}
	rawToken := RawToken{CurrentUser: make(map[string]interface{})}
	err = json.Unmarshal(byteResponse, &rawToken)
	if err != nil {
		return NewFdkTokenIssueError(err.Error())
	}
	o.SetAccessToken(rawToken)
	return nil
}
