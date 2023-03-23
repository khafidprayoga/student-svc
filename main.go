package main

import (
	"net/http"

	"buf.build/gen/go/khafidprayoga/student-svc/bufbuild/connect-go/student/v1/studentv1connect"
	"github.com/khafidprayoga/student-svc/svc"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	services := &svc.StudentServicesHandlerV1{}
	path, handler := studentv1connect.NewStudentServiceHandler(services)
	mux.Handle(path, handler)

	http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}
