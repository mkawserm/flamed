package context

import "net/http"

type GraphQLContext struct {
	URL        string
	Host       string
	RequestURI string
	Header     http.Header
	RemoteAddr string

	Data interface{}
}
