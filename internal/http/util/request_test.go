package util

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetHeaders(t *testing.T) {
	type args struct {
		r       *http.Request
		headers map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case nil",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				}(),
				headers: nil,
			},
		},
		{
			name: "case empty",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				}(),
				headers: map[string]string{},
			},
		},
		{
			name: "case normal",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				}(),
				headers: map[string]string{
					"Custom-Header-1": "custom_value_1",
					"Custom-Header-2": "custom_value_2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetHeaders(tt.args.r, tt.args.headers)
			for k, v := range tt.args.headers {
				assert.Equal(t, v, tt.args.r.Header.Get(k))
			}
		})
	}
}

func TestSetCookies(t *testing.T) {
	type args struct {
		r       *http.Request
		cookies []*http.Cookie
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case nil",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				}(),
				cookies: nil,
			},
		},
		{
			name: "case empty",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				}(),
				cookies: []*http.Cookie{},
			},
		},
		{
			name: "case normal",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				}(),
				cookies: []*http.Cookie{
					{
						Name:  "cookie_1",
						Value: "cookie-value-1",
					},
					{
						Name:  "cookie_2",
						Value: "cookie-value-2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetCookies(tt.args.r, tt.args.cookies)
			for _, c := range tt.args.cookies {
				cookie, err := tt.args.r.Cookie(c.Name)
				assert.Nil(t, err)
				assert.Equal(t, c, cookie)
			}
		})
	}
}
