package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

//RequestSigner holds request signer object
type RequestSigner struct {
	UpdatedReq        *http.Request
	RequestedDateTime string
	SignQuery         bool
}

//HEADERS_TO_IGNORE contains headers to be ignored while signing request
var HEADERS_TO_IGNORE = []string{
	"authorization",
	"connection",
	"x-amzn-trace-id",
	"user-agent",
	"expect",
	"presigned-expires",
	"range",
}

//HEADERS_TO_INCLUDE contains headers to be included while signing request
var HEADERS_TO_INCLUDE = []string{"host", "x-fp-.*"}

//NewRequestSigner return the request signer instance
func NewRequestSigner(req *http.Request, signQuery bool) *RequestSigner {
	return &RequestSigner{
		UpdatedReq:        req,
		RequestedDateTime: "",
		SignQuery:         signQuery,
	}
}

//GetUpdatedRequest return the updated HTTP request
func (r *RequestSigner) GetUpdatedRequest() *http.Request {
	return r.UpdatedReq
}

//GetRequestedDateTime reuse and returns the date time created
func (r *RequestSigner) GetRequestedDateTime() string {
	if r.RequestedDateTime != "" {
		return r.RequestedDateTime
	}
	return strings.Replace(strings.Replace(time.Now().UTC().Format(time.RFC3339), "-", "", -1), ":", "", -1)
}

//Sign performs request signing operation
func (r *RequestSigner) Sign() error {
	var (
		signature            string
		kCredentials         = "1234567"
		strToSign            string
		canonicalRequest     string
		canonicalQueryString string
		canonicalPath        string
		canonicalHeaders     string
		signedHeaders        string
		reqBodyHashed        string
	)

	r.UpdatedReq.Header["host"] = []string{r.UpdatedReq.Host}

	//Checking if request to be signed in query param or header
	if r.SignQuery {
		params := r.UpdatedReq.URL.Query()
		params["x-fp-date"] = []string{r.GetRequestedDateTime()}
		r.UpdatedReq.URL.RawQuery = params.Encode()
	} else {
		r.UpdatedReq.Header["x-fp-date"] = []string{r.GetRequestedDateTime()}
	}

	//generate signedHeaders
	for key := range r.UpdatedReq.Header {
		if !contains(HEADERS_TO_IGNORE, strings.ToLower(key), false) {
			if contains(HEADERS_TO_INCLUDE, strings.ToLower(key), true) {
				signedHeaders += fmt.Sprintf("%s;", strings.Trim(strings.ToLower(key), " "))
			}
		}
	}
	signedHeaders = strings.TrimSuffix(signedHeaders, ";")
	signedHeaders = sortHeaders(signedHeaders, ";")

	//generate canonical headers
	for key, val := range r.UpdatedReq.Header {
		if !contains(HEADERS_TO_IGNORE, strings.ToLower(key), false) {
			if contains(HEADERS_TO_INCLUDE, strings.ToLower(key), true) {
				canonicalHeaders += fmt.Sprintf("%s:%s\n", strings.ToLower(key), strings.Trim(val[0], " "))
			}
		}
	}
	canonicalHeaders = strings.TrimSuffix(canonicalHeaders, "\n")
	canonicalHeaders = sortHeaders(canonicalHeaders, "\n")

	// //generate canonical query string
	if strings.IndexAny(r.UpdatedReq.URL.String(), "?") > -1 {
		trimedURL := r.UpdatedReq.URL.String()[strings.IndexAny(r.UpdatedReq.URL.String(), "?")+1:]
		splittedQueryParams := strings.Split(trimedURL, "&")
		sort.SliceStable(splittedQueryParams, func(i, j int) bool {
			return splittedQueryParams[i] < splittedQueryParams[j]
		})
		//Converting query params to non-encoded string
		for i := 0; i < len(splittedQueryParams); i++ {
			val, err := url.QueryUnescape(splittedQueryParams[i])
			if err != nil {
				return errors.New("Cannot decode query param : " + err.Error())
			}
			splittedQueryParams[i] = val
		}
		canonicalQueryString = strings.Join(splittedQueryParams, "&")
	}

	//generate canonical path
	canonicalPath = r.UpdatedReq.URL.Path

	//generate req body hash in hex(sha256())
	reqBody, err := ioutil.ReadAll(r.UpdatedReq.Body)
	if err != nil {
		log.Println("Failed to read request body. Cannot sign the request")
		return errors.New("Cannot sign the request. Failed to read request body")
	}

	// Restore the request body to its original state
	r.UpdatedReq.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	reqBodyHashed = hash(reqBody)
	if reqBodyHashed == "" {
		log.Println("Cannot sign the request. Empty request body hash")
		return errors.New("Cannot sign the request. Empty request body hash")
	}

	//generate canonical request
	canonicalRequest = fmt.Sprintf("%s\n%s\n%s\n%s\n\n%s\n%s",
		r.UpdatedReq.Method,
		canonicalPath,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		reqBodyHashed)

	canonicalRequestHashed := hash([]byte(canonicalRequest))
	if canonicalRequestHashed == "" {
		log.Println("Cannot sign the request. Empty canonical request hash")
		return errors.New("Cannot sign the request. Empty canonical request hash")
	}

	//generate string to sign
	strToSign = fmt.Sprintf("%s\n%s", r.GetRequestedDateTime(), canonicalRequestHashed)

	//generate final signature
	h := hmac.New(sha256.New, []byte(kCredentials))
	_, err = h.Write([]byte(strToSign))
	if err != nil {
		log.Println("Cannot sign the request. Failed to write strToSign to hash")
		return errors.New("Cannot sign the request. Failed to write strToSign to hash")
	}
	sha := hex.EncodeToString(h.Sum(nil))
	signature = fmt.Sprintf("v1:%s", sha)
	if r.SignQuery {
		params := r.UpdatedReq.URL.Query()
		params["x-fp-signature"] = []string{signature}
		r.UpdatedReq.URL.RawQuery = params.Encode()
	} else {
		r.UpdatedReq.Header["x-fp-signature"] = []string{signature}
	}
	return nil
}

func hash(b []byte) string {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		log.Println("Failed to write request body to hash. err = ", err)
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// contains checks if a string is present in a slice
func contains(s []string, str string, isRegex bool) bool {
	for _, v := range s {
		if isRegex {
			r, _ := regexp.Compile(v)
			if !r.MatchString(str) {
				continue
			} else {
				return true
			}
		}
		if v == str {
			return true
		}
	}
	return false
}

func sortHeaders(header, cutset string) string {
	splittedHeaders := strings.Split(header, cutset)
	sort.SliceStable(splittedHeaders, func(i, j int) bool {
		return splittedHeaders[i] < splittedHeaders[j]
	})
	return strings.Join(splittedHeaders, cutset)
}
