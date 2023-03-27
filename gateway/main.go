package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khafidprayoga/student-svc/common/config"
	"github.com/khafidprayoga/student-svc/common/constant"
	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		mux         = runtime.NewServeMux()
		log         = config.GetZapLogger()
	)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Info("registering grpc service for HTTP endpoint")
	errSetupGW := studentv2.RegisterStudentServiceHandlerFromEndpoint(
		ctx,
		mux, constant.GRPCAddress,
		opts,
	)

	if errSetupGW != nil {
		panic(errSetupGW)
	}

	startProxyMsg := fmt.Sprintf(
		"serving proxy on %v ",
		constant.GatewayAddress,
	)

	log.Info(startProxyMsg)
	serveErr := http.ListenAndServe(constant.GatewayAddress, mux)
	if serveErr != nil {
		panic(serveErr)
	}
}
