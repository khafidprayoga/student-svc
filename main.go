package main

import (
	"fmt"
	"net"

	"buf.build/gen/go/khafidprayoga/student-svc/grpc/go/student/v2/studentv2grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/khafidprayoga/student-svc/common/config"
	"github.com/khafidprayoga/student-svc/common/constant"
	"github.com/khafidprayoga/student-svc/common/data"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/khafidprayoga/student-svc/svc"
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

	studentv2grpc.RegisterStudentServiceServer(s, handler)
	reflection.Register(s)

	if e := s.Serve(listen); e != nil {
		panic(e)
	}
}
