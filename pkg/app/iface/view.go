package iface

import "net/http"

type IView interface {
	GetHTTPHandler() http.HandlerFunc
}
