package main

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/khafidprayoga/student-svc/common/config"
	"github.com/khafidprayoga/student-svc/common/constant"
	"github.com/khafidprayoga/student-svc/common/data"
	"github.com/khafidprayoga/student-svc/gen/student/v2"
	"github.com/khafidprayoga/student-svc/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var (
		log       = config.GetZapLogger()
		studentDB = data.New()
	)

	listen, err := net.Listen("tcp", constant.GRPCAddress)
	if err != nil {
		panic(err)
	}

	bootMsg := fmt.Sprintf(
		"student services running on %v",
		constant.GRPCAddress,
	)

	log.Info(bootMsg)

	handler := &svc.StudentServiceServerImpl{
		Log:  log,
		Data: studentDB,
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
