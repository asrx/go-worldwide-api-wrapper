package go_worldwide_api_wrapper

import (
	"strings"

	"github.com/idoubi/goz"
)

const api_basic_url = "https://label-api.rapiddeals.com"
const ContentType = "application/json"

func getHeader(id, key string) map[string]interface{} {
	return map[string]interface{}{
		"api-id":       id,
		"api-key":      key,
		"Content-Type": ContentType,
	}
}

func getRequestUrl(targetUrl string) string {
	return strings.TrimSuffix(api_basic_url, "/") + "/" + strings.TrimPrefix(targetUrl, "/")
}

func NewHttpClient(method, url string) (*goz.Response, error) {

	client := goz.NewClient()
	return client.Request(method, api_basic_url, goz.Options{
		Timeout:    3000,
		Headers:    nil,
		FormParams: nil,
	})
}
