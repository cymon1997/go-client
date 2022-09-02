package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cymon1997/go-client/http/util"
	"github.com/cymon1997/go-client/internal/utils"
	"github.com/google/go-querystring/query"
)

type Request interface {
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

type requestImpl struct {
	client  *http.Client
	baseURL string
	body    []byte
	headers map[string]string
	cookies []*http.Cookie
}

func NewRequest(client *http.Client, baseURL string, headers map[string]string) Request {
	return &requestImpl{
		client:  client,
		baseURL: baseURL,
		headers: headers,
	}
}

func (r *requestImpl) WithHeaders(headers map[string]string) Request {
	r.headers = utils.CombineMapString(r.headers, headers, utils.MergeReplace)
	return r
}

func (r *requestImpl) WithCookies(cookies []*http.Cookie) Request {
	r.cookies = cookies
	return r
}

// Get used for retrieve a resource
func (r *requestImpl) Get(ctx context.Context, endpoint string, params interface{}) (*http.Response, error) {
	v, _ := query.Values(params)
	fmt.Println("DEBUG: ", v.Encode())
	return r.exec(ctx, http.MethodGet, fmt.Sprintf("%s?%s", endpoint, v.Encode()))
}

// Post used for create a resource
func (r *requestImpl) Post(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return r.execBody(ctx, http.MethodPost, endpoint, body)
}

// PostRaw is raw version of Post, usually used for upload raw file
func (r *requestImpl) PostRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error) {
	return r.execRaw(ctx, http.MethodPost, endpoint, raw)
}

// Put used for update & create a resource
func (r *requestImpl) Put(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return r.execBody(ctx, http.MethodPut, endpoint, body)
}

// PutRaw is raw version of Put, usually used for upload raw file
func (r *requestImpl) PutRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error) {
	return r.execRaw(ctx, http.MethodPut, endpoint, raw)
}

// Patch used for update a resource
func (r *requestImpl) Patch(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	return r.execBody(ctx, http.MethodPatch, endpoint, body)
}

// PatchRaw is raw version of Patch, usually used for upload raw file
func (r *requestImpl) PatchRaw(ctx context.Context, endpoint string, raw []byte) (*http.Response, error) {
	return r.execRaw(ctx, http.MethodPatch, endpoint, raw)
}

// Delete used for delete a resource
func (r *requestImpl) Delete(ctx context.Context, endpoint string) (*http.Response, error) {
	return r.exec(ctx, http.MethodDelete, endpoint)
}

func (r *requestImpl) execBody(ctx context.Context, method, endpoint string, body interface{}) (*http.Response, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return r.execRaw(ctx, method, endpoint, raw)
}

func (r *requestImpl) execRaw(ctx context.Context, method, endpoint string, raw []byte) (*http.Response, error) {
	r.body = raw
	return r.exec(ctx, method, endpoint)
}

func (r *requestImpl) exec(ctx context.Context, method, uri string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprint(r.baseURL, uri), bytes.NewBuffer(r.body))
	if err != nil {
		return nil, err
	}
	util.SetHeaders(req, r.headers)
	util.SetCookies(req, r.cookies)
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
