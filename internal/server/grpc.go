// Package server hosts the network adapters. Generated Envoy-compatible handlers plug in here.
package server

import (
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct{ server *grpc.Server }

func NewGRPCServer() *GRPCServer { return &GRPCServer{server: grpc.NewServer()} }
func (s *GRPCServer) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.server.Serve(l)
}
func (s *GRPCServer) GracefulStop() { s.server.GracefulStop() }
