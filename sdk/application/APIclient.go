package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gofynd/fdk-client-golang/sdk/common"
	"github.com/gofynd/fdk-client-golang/sdk/common/utils"
	"github.com/google/go-querystring/query"
	"moul.io/http2curl"
)

//RawRequest holds request body data
type RawRequest struct {
	BaseURL     string
	APIURL      string
	Method      string
	Headers     map[string]string
	QueryParams interface{}
	Body        map[string]interface{}
}

//NewRequest process rew request body
func NewRequest(conf *AppConfig, method, apiURL string, headers map[string]string, query interface{}, body map[string]interface{}) *RawRequest {
	token := utils.EncodeToBase64(fmt.Sprintf("%s:%s", conf.ApplicationID, conf.ApplicationToken))
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
	}
}

//Execute performs API call
func (r *RawRequest) Execute() ([]byte, error) {
	req, err := processHTTPRequest(r)
	if err != nil {
		return []byte{}, common.NewFDKError(err.Error())
	}
	reqSigner := common.NewRequestSigner(req)
	err = reqSigner.Sign()
	if err != nil {
		return []byte{}, common.NewFDKError(err.Error())
	}
	if len(req.Header["x-fp-signature"]) > 0 {
		log.Println("Request signing successfull...signature = ", req.Header["x-fp-signature"])
	}
	command, _ := http2curl.GetCurlCommand(req)
	log.Println("curl = ", command)
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
		reqBodyJSON, err := json.Marshal(r.Body)
		if err != nil {
			return &http.Request{}, err
		}
		payload = strings.NewReader(string(reqBodyJSON))
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
