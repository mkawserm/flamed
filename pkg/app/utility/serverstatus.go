package utility

import (
	"sync"
)

var (
	serverStatusOnce sync.Once
	serverStatusIns  *ServerStatus
)

type ServerStatus struct {
	mHTTPServer bool
	mGRPCServer bool
	mRAFTServer bool

	mMutex sync.Mutex
}

func (s *ServerStatus) HTTPServer() bool {
	s.mMutex.Lock()
	defer s.mMutex.Unlock()

	return s.mHTTPServer
}

func (s *ServerStatus) SetHTTPServer(b bool) {
	s.mMutex.Lock()
	defer s.mMutex.Unlock()

	s.mHTTPServer = b
}

func (s *ServerStatus) GRPCServer() bool {
	s.mMutex.Lock()
	defer s.mMutex.Unlock()

	return s.mGRPCServer
}

func (s *ServerStatus) SetGRPCServer(b bool) {
	s.mMutex.Lock()
	defer s.mMutex.Unlock()

	s.mGRPCServer = b
}

func (s *ServerStatus) RAFTServer() bool {
	s.mMutex.Lock()
	defer s.mMutex.Unlock()

	return s.mRAFTServer
}

func (s *ServerStatus) SetRAFTServer(b bool) {
	s.mMutex.Lock()
	defer s.mMutex.Unlock()

	s.mRAFTServer = b
}

func GetServerStatus() *ServerStatus {
	return serverStatusIns
}

func init() {
	serverStatusOnce.Do(func() {
		serverStatusIns = &ServerStatus{}
	})
}
