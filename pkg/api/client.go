package api

import (
	"context"
	"net/http"
	"time"
)

type Client interface {
	SetBaseHeaders(headers map[string]string)
	WithHeaders(headers map[string]string) Request
	WithCookies(cookies []*http.Cookie) Request

	Get(ctx context.Context, endpoint string, params interface{}) (*http.Response, error)
	Post(ctx context.Context, endpoint string, body interface{}) (*http.Response, error)
	PostRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error)
	Put(ctx context.Context, endpoint string, body interface{}) (*http.Response, error)
	PutRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error)
	Patch(ctx context.Context, endpoint string, body interface{}) (*http.Response, error)
	PatchRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error)
	Delete(ctx context.Context, endpoint string) (*http.Response, error)
}

type clientImpl struct {
	client      *http.Client
	baseURL     string
	baseHeaders map[string]string
}

func New(cfg Config) Client {
	return &clientImpl{
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Millisecond,
		},
		baseURL: cfg.Host,
	}
}

func (c *clientImpl) SetBaseHeaders(headers map[string]string) {
	c.baseHeaders = headers
}

func (c *clientImpl) WithHeaders(headers map[string]string) Request {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).WithHeaders(headers)
}

func (c *clientImpl) WithCookies(cookies []*http.Cookie) Request {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).WithCookies(cookies)
}

func (c *clientImpl) Get(ctx context.Context, endpoint string, params interface{}) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).Get(ctx, endpoint, params)
}

func (c *clientImpl) Post(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).Post(ctx, endpoint, body)
}

func (c *clientImpl) PostRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).PostRaw(ctx, endpoint, raw)
}

func (c *clientImpl) Put(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).Put(ctx, endpoint, body)
}

func (c *clientImpl) PutRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).PutRaw(ctx, endpoint, raw)
}

func (c *clientImpl) Patch(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).Patch(ctx, endpoint, body)
}

func (c *clientImpl) PatchRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).PatchRaw(ctx, endpoint, raw)
}

func (c *clientImpl) Delete(ctx context.Context, endpoint string) (*http.Response, error) {
	return NewRequest(c.client, c.baseURL, c.baseHeaders).Delete(ctx, endpoint)
}
