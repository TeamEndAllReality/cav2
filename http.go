package cav2

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//GetHTTPResponse gets response from URL (Wraps GetAuthenicatedResponse)
func GetHTTPResponse(method, url string, b []byte) (*http.Response, error) {
	return GetAuthenticatedReponse(method, OAuthToken, url, b)
}

//GetAuthenticatedReponse wraps an http client with an authenication token
func GetAuthenticatedReponse(method, token, url string, b []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if len(token) > 0 && strings.HasPrefix(url, APIEndpoint) {
		req.Header.Add("AuthenticationToken", token)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return resp, errors.New(resp.Status)
	}

	return resp, err
}

//DoHTTPRequest does a basic http request
func DoHTTPRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, bytes.NewReader(EMPTY))
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return resp, errors.New(resp.Status)
	}
	return resp, err
}

//ResponseToBytes Converts an http.Response to bytes
func ResponseToBytes(response *http.Response) []byte {
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	return bodyBytes
}

//ResponseToString Converts an http.Response to a string
func ResponseToString(response *http.Response) string {
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	return bodyString
}
