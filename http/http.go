package httpclient
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var JsonHeaders = map[string]string{
	"Content-Type": "application/json",
	"Accept":       "application/json",
}

func Request[T any](method, url string, data any, headers map[string]string) (result T, respBody []byte, err error) {
	var client = otelhttp.DefaultClient
	var body io.Reader
	var jsonData []byte

	if data != nil {
		jsonData, err = json.Marshal(data)
		if err != nil {
			return result, nil, fmt.Errorf("marshal request error %s", err.Error())
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return result, nil, fmt.Errorf("create request error %s", err.Error())
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, nil, fmt.Errorf("send request error %s", err.Error())
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return result, respBody, fmt.Errorf("read response error %s", err.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return result, nil, fmt.Errorf("HTTP request failed with status code %d: %s", resp.StatusCode, respBody)
	}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return result, respBody, fmt.Errorf("decode response error %s", err.Error())
	}
	return
}
