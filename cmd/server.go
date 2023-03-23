package main

import (
	"fmt"
	"net"

	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
	"github.com/khafidprayoga/student-svc/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = ":5080"

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Printf("student services running on %v\n", address)
	handler := &svc.StudentServiceServerImpl{}
	s := grpc.NewServer()
	studentv2.RegisterStudentServiceServer(s, handler)
	reflection.Register(s)

	if e := s.Serve(listen); e != nil {
		panic(e)
	}
}
