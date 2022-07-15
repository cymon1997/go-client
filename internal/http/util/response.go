package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IsStatusOK check if response code is 2xx
func IsStatusOK(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// ParseBody parse request body to dest struct
func ParseBody(resp *http.Response, dest interface{}) error {
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.Unmarshal(raw, &dest)
	if err != nil {
		return err
	}
	return nil
}
