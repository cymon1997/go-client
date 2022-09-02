package util

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestIsStatusOK(t *testing.T) {
	t.Run("http.util.Response", func(t *testing.T) {
		defer gock.OffAll()

		client := http.Client{}
		request, _ := http.NewRequest(
			http.MethodGet, "http://localhost:8000/get", nil)

		// Case 1xx
		gock.New("http://localhost:8000").
			Get("/get").
			Reply(199)

		resp, err := client.Do(request)
		assert.Nil(t, err)
		assert.False(t, IsStatusOK(resp))

		// Case 2xx
		gock.New("http://localhost:8000").
			Get("/get").
			Reply(200)

		resp, err = client.Do(request)
		assert.Nil(t, err)
		assert.True(t, IsStatusOK(resp))

		// Case 3xx
		gock.New("http://localhost:8000").
			Get("/get").
			Reply(300)

		resp, err = client.Do(request)
		assert.Nil(t, err)
		assert.False(t, IsStatusOK(resp))
	})
}

func TestParseBody(t *testing.T) {
	t.Run("http.util.Response", func(t *testing.T) {
		defer gock.OffAll()

		type Response struct {
			Data string `json:"data"`
		}

		client := http.Client{}
		request, _ := http.NewRequest(
			http.MethodGet, "http://localhost:8000/get", nil)
		var response Response

		// Case error
		gock.New("http://localhost:8000").
			Get("/get").
			Reply(400).
			Body(bytes.NewReader([]byte(`{`)))

		resp, err := client.Do(request)
		assert.Nil(t, err)
		err = ParseResponseBody(resp, &response)
		assert.Error(t, err)

		// Case success
		gock.New("http://localhost:8000").
			Get("/get").
			Reply(200).
			Body(bytes.NewReader([]byte(`{"data":"some_data"}`)))

		resp, err = client.Do(request)
		assert.Nil(t, err)
		err = ParseResponseBody(resp, &response)
		assert.Nil(t, err)
		assert.Equal(t, "some_data", response.Data)
	})
}
