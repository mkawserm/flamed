package server

import (
	iface2 "github.com/mkawserm/flamed/pkg/app/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"sync"
)

type GRPCServer struct {
	mMutex    sync.Mutex
	mServer   *grpc.Server
	mServices []iface2.IGRPCService
}

func (g *GRPCServer) AddService(service iface2.IGRPCService) {
	g.mMutex.Lock()
	defer g.mMutex.Unlock()

	g.mServices = append(g.mServices, service)
}

func (g *GRPCServer) GracefulStop() {
	g.mMutex.Lock()
	defer g.mMutex.Unlock()

	if g.mServer != nil {
		g.mServer.GracefulStop()
		g.mServer = nil
	}
}

func (g *GRPCServer) Start(address string,
	enableTLS bool,
	certFile string,
	keyFile string) error {
	g.mMutex.Lock()
	if g.mServer != nil {
		g.mMutex.Unlock()
		return nil
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if enableTLS {
		// Create the TLS transportCredentials
		transportCredentials, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return err
		}
		// Create an array of gRPC options with the transportCredentials
		opts := []grpc.ServerOption{grpc.Creds(transportCredentials)}
		g.mServer = grpc.NewServer(opts...)

		for _, v := range g.mServices {
			v.RegisterGRPCService(g.mServer)
		}

		logger.L("grpc::server").Info("grpc server with tls started @ " + address)
		g.mMutex.Unlock()

		return g.mServer.Serve(lis)
	} else {
		g.mServer = grpc.NewServer()
		for _, v := range g.mServices {
			v.RegisterGRPCService(g.mServer)
		}
		logger.L("grpc::server").Info("grpc server started @ " + address)
		g.mMutex.Unlock()

		return g.mServer.Serve(lis)
	}
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		mServices: make([]iface2.IGRPCService, 0, 50),
	}
}
