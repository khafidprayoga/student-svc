package main

import (
	"fmt"
	"net"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/khafidprayoga/student-svc/cmd/config"
	"github.com/khafidprayoga/student-svc/gen/student/v2"
	"github.com/khafidprayoga/student-svc/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = ":5080"

func main() {
	log := config.GetZapLogger()
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	bootMsg := fmt.Sprintf("student services running on %v", address)
	log.Info(bootMsg)

	handler := &svc.StudentServiceServerImpl{
		Log: log,
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	studentv2.RegisterStudentServiceServer(s, handler)
	reflection.Register(s)

	if e := s.Serve(listen); e != nil {
		panic(e)
	}
}
