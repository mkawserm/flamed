package server

import (
	"context"
	"github.com/mkawserm/flamed/pkg/app/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"net/http"
)

type HTTPServer struct {
	mServerMux  *http.ServeMux
	mHTTPServer *http.Server
}

func (h *HTTPServer) AddView(pattern string, view iface.IHTTPView) {
	h.mServerMux.HandleFunc(pattern, view.GetHTTPHandler())
}

func (h *HTTPServer) AddHandlerFunc(pattern string, handler http.HandlerFunc) {
	h.mServerMux.HandleFunc(pattern, handler)
}

func (h *HTTPServer) Start(address string, enableTLS bool, certFile string, keyFile string) error {
	h.mHTTPServer.Addr = address

	if enableTLS {
		logger.L("http::server").Info("http server with tls started @ " + address)
		return h.mHTTPServer.ListenAndServeTLS(certFile, keyFile)
	} else {
		logger.L("http::server").Info("http server started @ " + address)
		return h.mHTTPServer.ListenAndServe()
	}
}

func (h *HTTPServer) Shutdown(ctx context.Context) error {
	return h.mHTTPServer.Shutdown(ctx)
}

func NewHTTPServer() *HTTPServer {
	s := &HTTPServer{
		mServerMux:  &http.ServeMux{},
		mHTTPServer: &http.Server{},
	}
	s.mHTTPServer.Handler = s.mServerMux

	return s
}
