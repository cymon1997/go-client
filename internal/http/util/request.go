package util

import (
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
