package examples

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cymon1997/go-client/internal/http/util"
	"github.com/cymon1997/go-client/pkg/api"
)

func examples() {
	cfg := api.Config{
		Host:    "http://localhost:8000",
		Timeout: 3000,
	}

	// Init client
	client := api.New(cfg)

	// Set mandatory headers that will be applied to all requests
	client.SetBaseHeaders(map[string]string{
		"X-API-Key":     "sample_api_key",
		"Authorization": "sample_auth",
	})

	ctx := context.Background()

	// GET request (with no params)
	resp, err := client.Get(ctx, "/get", nil)
	if err != nil {
		log.Println("error: ", err)
		return
	}
	// Check if status is 2xx
	if !util.IsStatusOK(resp) {
		log.Println("status not OK")
		return
	}
	// Automatic parse response with safe close
	type Response struct {
		Message int         `json:"message"`
		Data    interface{} `json:"data"`
	}
	var response Response
	err = util.ParseBody(resp, &response)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	// GET request (using auto params)
	type Params struct {
		Query string `json:"query"`
		Page  int    `json:"page"`
	}
	resp, err = client.Get(ctx, "/search", Params{
		Query: "sample search query",
		Page:  1,
	})
	// Result param: /search?query=sample%20search%20query&page=1

	// POST request using object as request body
	type Body struct {
		Data string `json:"data"`
	}
	resp, err = client.Post(ctx, "/post", Body{
		Data: "sample_data",
	})
	// POST request using raw encoded byte
	raw, _ := json.Marshal(Body{
		Data: "sample_data",
	})
	resp, err = client.PostRaw(ctx, "/post", raw)

	// PUT request using object as request body
	resp, err = client.Put(ctx, "/put", Body{
		Data: "sample_data",
	})
	// PUT request using raw encoded byte
	resp, err = client.PutRaw(ctx, "/put", raw)

	// PATCH request using object as request body
	resp, err = client.Patch(ctx, "/patch", Body{
		Data: "sample_data",
	})
	// PATCH request using raw encoded byte
	resp, err = client.PatchRaw(ctx, "/patch", raw)

	// DELETE request using object as request body
	resp, err = client.Delete(ctx, "/delete/1")

	// Request with custom headers & cookies
	resp, err = client.WithHeaders(map[string]string{
		"Custom-Header": "sample_header_value",
	}).WithCookies([]*http.Cookie{
		{
			Name:   "cookie_name",
			Value:  "sample_cookie_value",
			Path:   "/",
			MaxAge: 300,
		},
	}).Get(ctx, "/get", nil)
}
