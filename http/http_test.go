package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}

	tests := []struct {
		name           string
		method         string
		url            string
		data           any
		headers        map[string]string
		expectedResult response
		expectedError  string
		statusCode     int
		responseBody   string
	}{
		{
			name:           "GET request success",
			method:         http.MethodGet,
			url:            "/test_success",
			data:           nil,
			headers:        JsonHeaders,
			expectedResult: response{Message: "success"},
			expectedError:  "",
			statusCode:     http.StatusOK,
			responseBody:   `{"message":"success"}`,
		},
		{
			name:           "POST request with data",
			method:         http.MethodPost,
			url:            "/test_create",
			data:           map[string]string{"key": "value"},
			headers:        JsonHeaders,
			expectedResult: response{Message: "created"},
			expectedError:  "",
			statusCode:     http.StatusCreated,
			responseBody:   `{"message":"created"}`,
		},
		{
			name:           "Request failure",
			method:         http.MethodGet,
			url:            "/test_fail",
			data:           nil,
			headers:        JsonHeaders,
			expectedResult: response{},
			expectedError:  "HTTP request failed with status code 500: Internal Server Error",
			statusCode:     http.StatusInternalServerError,
			responseBody:   `Internal Server Error`,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, tt := range tests {
			if r.URL.Path == tt.url && r.Method == tt.method {
				w.WriteHeader(tt.statusCode)
				_, err := w.Write([]byte(tt.responseBody))
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			resp, respBody, err := Request[response](ctx, tt.method, server.URL+tt.url, tt.data, tt.headers)

			if tt.expectedError != "" {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, resp)
				assert.Equal(t, []byte(tt.responseBody), respBody)
				assert.NotNil(t, resp)
			}
		})
	}
}
