package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

type RequestSigner struct {
	UpdatedReq        *http.Request
	RequestedDateTime string
}

var HEADERS_TO_IGNORE = []string{
	"authorization",
	"connection",
	"x-amzn-trace-id",
	"user-agent",
	"expect",
	"presigned-expires",
	"range",
}

var HEADERS_TO_INCLUDE = []string{"host", "x-fp-.*"}

func NewRequestSigner(req *http.Request) *RequestSigner {
	return &RequestSigner{
		UpdatedReq:        req,
		RequestedDateTime: "",
	}
}

func (r *RequestSigner) GetUpdatedRequest() *http.Request {
	return r.UpdatedReq
}

func (r *RequestSigner) GetRequestedDateTime() string {
	if r.RequestedDateTime != "" {
		return r.RequestedDateTime
	}
	return strings.Replace(strings.Replace(time.Now().UTC().Format(time.RFC3339), "-", "", -1), ":", "", -1)
}

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
	r.UpdatedReq.Header["x-fp-date"] = []string{r.GetRequestedDateTime()}

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
	r.UpdatedReq.Header["x-fp-signature"] = []string{signature}
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
