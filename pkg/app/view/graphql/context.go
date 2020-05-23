package graphql

import "net/http"

type Context struct {
	URL        string
	Host       string
	RequestURI string
	Header     http.Header
	RemoteAddr string
}
