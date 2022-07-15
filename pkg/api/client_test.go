package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			name: "case normal",
			args: args{
				cfg: Config{
					Host:    "http://localhost:8000",
					Timeout: 3000,
				},
			},
			want: &clientImpl{
				client: &http.Client{
					Timeout: time.Duration(3000) * time.Millisecond,
				},
				baseURL: "http://localhost:8000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clientImpl_SetBaseHeaders(t *testing.T) {
	type fields struct {
		client      *http.Client
		baseURL     string
		baseHeaders map[string]string
	}
	type args struct {
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case normal",
			fields: fields{
				baseHeaders: nil,
			},
			args: args{
				headers: map[string]string{
					"X-BASE-HEADER": "some_value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &clientImpl{
				client:      tt.fields.client,
				baseURL:     tt.fields.baseURL,
				baseHeaders: tt.fields.baseHeaders,
			}
			c.SetBaseHeaders(tt.args.headers)
			assert.Equal(t, tt.args.headers, c.baseHeaders)
		})
	}
}

func Test_clientImpl_Get(t *testing.T) {
	t.Run("client.Get", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		params := struct {
			Query string `url:"query"`
			Page  int    `url:"page"`
		}{
			Query: "sample query",
			Page:  1,
		}

		gock.New("http://localhost:8000").
			Get("/search").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			MatchParams(map[string]string{
				"query": "sample query",
				"page":  "1",
			}).Reply(200)

		got, err := c.Get(context.Background(), "/search", params)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Get("/search").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			MatchParams(map[string]string{
				"query": "sample query",
				"page":  "1",
			}).Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).Get(context.Background(), "/search", params)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_Post(t *testing.T) {
	t.Run("client.Post", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		body := struct {
			Data string `json:"data"`
		}{
			Data: "sample data",
		}

		raw, _ := json.Marshal(body)
		gock.New("http://localhost:8000").
			Post("/post").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err := c.Post(context.Background(), "/post", body)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Post("/post").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).Post(context.Background(), "/post", body)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_PostRaw(t *testing.T) {
	t.Run("client.PostRaw", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		body := struct {
			Data string `json:"data"`
		}{
			Data: "sample data",
		}

		raw, _ := json.Marshal(body)
		gock.New("http://localhost:8000").
			Post("/post").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err := c.PostRaw(context.Background(), "/post", raw)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Post("/post").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).PostRaw(context.Background(), "/post", raw)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_Put(t *testing.T) {
	t.Run("client.Put", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		body := struct {
			Data string `json:"data"`
		}{
			Data: "sample data",
		}

		raw, _ := json.Marshal(body)
		gock.New("http://localhost:8000").
			Put("/put").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err := c.Put(context.Background(), "/put", body)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Put("/put").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).Put(context.Background(), "/put", body)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_PutRaw(t *testing.T) {
	t.Run("client.PutRaw", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		body := struct {
			Data string `json:"data"`
		}{
			Data: "sample data",
		}

		raw, _ := json.Marshal(body)
		gock.New("http://localhost:8000").
			Put("/put").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err := c.PutRaw(context.Background(), "/put", raw)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Put("/put").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).PutRaw(context.Background(), "/put", raw)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_Patch(t *testing.T) {
	t.Run("client.Patch", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		body := struct {
			Data string `json:"data"`
		}{
			Data: "sample data",
		}

		raw, _ := json.Marshal(body)
		gock.New("http://localhost:8000").
			Patch("/patch").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err := c.Patch(context.Background(), "/patch", body)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Patch("/patch").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).Patch(context.Background(), "/patch", body)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_PatchRaw(t *testing.T) {
	t.Run("client.PatchRaw", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		body := struct {
			Data string `json:"data"`
		}{
			Data: "sample data",
		}

		raw, _ := json.Marshal(body)
		gock.New("http://localhost:8000").
			Patch("/patch").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err := c.PatchRaw(context.Background(), "/patch", raw)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Patch("/patch").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Body(bytes.NewReader(raw)).
			Reply(200)

		got, err = c.WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).PatchRaw(context.Background(), "/patch", raw)
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}

func Test_clientImpl_Delete(t *testing.T) {
	t.Run("client.Delete", func(t *testing.T) {
		defer gock.OffAll()

		c := &clientImpl{
			client: &http.Client{
				Timeout: 3000 * time.Millisecond,
			},
			baseURL: "http://localhost:8000",
			baseHeaders: map[string]string{
				"X-API-Key": "sample_api_key",
			},
		}

		gock.New("http://localhost:8000").
			MatchHeaders(map[string]string{
				"X-API-Key": "sample_api_key",
			}).
			Delete("/delete/1").
			Reply(200)

		got, err := c.Delete(context.Background(), "/delete/1")
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)

		gock.New("http://localhost:8000").
			Delete("/delete/1").
			MatchHeaders(map[string]string{
				"X-API-Key":       "sample_api_key",
				"Custom-Header-1": "custom_key_1",
				"Custom-Header-2": "custom_key_2",
			}).
			Reply(200)

		got, err = c.WithCookies([]*http.Cookie{
			{
				Name:  "sample_cookie",
				Value: "cookie_value",
				Path:  "/",
			},
		}).WithHeaders(map[string]string{
			"Custom-Header-1": "custom_key_1",
			"Custom-Header-2": "custom_key_2",
		}).
			Delete(context.Background(), "/delete/1")
		assert.Nil(t, err)
		assert.Equal(t, 200, got.StatusCode)
	})
}
