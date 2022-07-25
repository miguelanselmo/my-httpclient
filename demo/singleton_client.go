package demo

import (
	"net/http"
	"time"

	"github.com/miguelanselmo/my-httpclient/httpc"
	"github.com/miguelanselmo/my-httpclient/mime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() httpc.Client {
	headers := make(http.Header)
	headers.Set(mime.HeaderContentType, mime.ContentTypeJson)

	client := httpc.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Fedes-Computer").
		Build()
	return client
}

func getHttpClientH(headers http.Header) httpc.Client {
	client := httpc.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Fedes-Computer").
		Build()
	return client
}
