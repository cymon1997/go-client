package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SetHeaders(r *http.Request, headers map[string]string) {
	for key, val := range headers {
		r.Header.Set(key, val)
	}
}

func SetCookies(r *http.Request, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}
}

// ParseRequestBody parse request body to dest struct
func ParseRequestBody(req *http.Request, dest interface{}) error {
	raw, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	err = json.Unmarshal(raw, &dest)
	if err != nil {
		return err
	}
	return nil
}
