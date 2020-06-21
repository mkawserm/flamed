package iface

import "google.golang.org/grpc"

type IGRPCService interface {
	RegisterGRPCService(server *grpc.Server)
}
