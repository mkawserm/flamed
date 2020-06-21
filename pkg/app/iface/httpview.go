package iface

import "net/http"

type IHTTPView interface {
	GetHTTPHandler() http.HandlerFunc
}
