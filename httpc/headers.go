package httpc

import (
	"net/http"

	"github.com/miguelanselmo/my-httpclient/mime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers from the HTTP client instance:
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers from the current request:
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Set User-Agent if it is defined and not there yet:
	if c.builder.userAgent != "" {
		if result.Get(mime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(mime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}
