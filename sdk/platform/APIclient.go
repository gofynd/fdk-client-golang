package platform

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gofynd/fdk-client-golang/sdk/common"
	"github.com/gofynd/fdk-client-golang/sdk/common/utils"
	"github.com/google/go-querystring/query"
)

//RawRequest holds request body data
type RawRequest struct {
	BaseURL     string
	APIURL      string
	Method      string
	Headers     map[string]string
	QueryParams interface{}
	Body        map[string]interface{}
	EncodeData  bool
}

//NewRequest prepares rew request body
func NewRequest(conf *PlatformConfig, method, apiURL string, headers map[string]string, query interface{}, body map[string]interface{}) *RawRequest {
	//GET token from oAuthClient
	token := conf.GetAccessToken()
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["authorization"] = "Bearer " + token
	return &RawRequest{
		BaseURL:     conf.Domain,
		APIURL:      apiURL,
		Method:      strings.ToUpper(method),
		Headers:     headers,
		QueryParams: query,
		Body:        body,
		EncodeData:  false,
	}
}

//Execute performs API call
func (r *RawRequest) Execute() ([]byte, error) {
	req, err := processHTTPRequest(r)
	if err != nil {
		return []byte{}, common.NewFDKError(err.Error())
	}
	reqSigner := common.NewRequestSigner(req, false)
	err = reqSigner.Sign()
	if err != nil {
		return []byte{}, common.NewFDKError(err.Error())
	}
	// if len(req.Header["x-fp-signature"]) > 0 {
	// log.Println("Request signing successfull...signature = ", req.Header["x-fp-signature"])
	// }
	// command, _ := http2curl.GetCurlCommand(req)
	// log.Println("curl = ", command)
	// timeout := time.Duration(15 * time.Second)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, common.NewFDKError(err.Error())
	}
	// read all response body
	data, err := processHTTPResponse(res)
	if err != nil {
		return []byte{}, common.NewFDKError(err.Error())
	}
	return data, nil
}

func processHTTPRequest(r *RawRequest) (*http.Request, error) {
	params := url.Values{}
	payload := &strings.Reader{}
	if r.APIURL == "" {
		return &http.Request{}, errors.New("No URL present in request, unable to sign request")
	}
	u, err := url.Parse(r.BaseURL)
	if err != nil {
		return &http.Request{}, err
	}
	u.Path = path.Join(u.Path, r.APIURL)
	r.APIURL = u.String()
	if r.Method == "GET" && r.QueryParams != nil {
		params, err = query.Values(r.QueryParams)
		// for k, v := range r.QueryParams {
		// 	params.Add(k, v)
		// 	// payload = strings.NewReader(params.Encode())
		// }
	}
	if r.Method != "GET" && r.Body != nil && len(r.Body) > 0 {
		payload, err = setRequestBody(r)
		if err != nil {
			return &http.Request{}, errors.New("Failed to set request body, unable to sign request : " + err.Error())
		}
	}
	req, err := http.NewRequest(r.Method, r.APIURL, payload)
	if err != nil {
		return &http.Request{}, err
	}
	//Setting headers
	for k, v := range r.Headers {
		req.Header[k] = []string{v}
	}
	//Setting query params
	req.URL.RawQuery = params.Encode()
	return req, nil
}

func processHTTPResponse(res *http.Response) ([]byte, error) {
	var errResp *common.FDKError
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	// log.Println("data", string(data))
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return []byte{}, err
		}
		return []byte{}, errResp
	}
	return data, nil
}

func setRequestBody(r *RawRequest) (*strings.Reader, error) {
	payload := &strings.Reader{}
	reqBodyJSON, err := json.Marshal(r.Body)
	if err != nil {
		return &strings.Reader{}, err
	}
	if r.EncodeData {
		reqBodyMap := make(map[string]string)
		err = json.Unmarshal(reqBodyJSON, &reqBodyMap)
		if err != nil {
			return &strings.Reader{}, err
		}
		payload = strings.NewReader(utils.MapToURLEncodedString(reqBodyMap))
	} else {
		payload = strings.NewReader(string(reqBodyJSON))
	}
	return payload, nil
}
