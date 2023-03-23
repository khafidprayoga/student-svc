package svc

import (
	"context"
	"errors"

	"buf.build/gen/go/khafidprayoga/student-svc/bufbuild/connect-go/student/v1/studentv1connect"
	v1 "buf.build/gen/go/khafidprayoga/student-svc/protocolbuffers/go/student/v1"
	connect_go "github.com/bufbuild/connect-go"
)

type StudentServicesHandlerV1 struct {
	studentv1connect.UnimplementedStudentServiceHandler
}

func (*StudentServicesHandlerV1) CreateStudent(
	ctx context.Context, req *connect_go.Request[v1.CreateStudentRequest]) (
	res *connect_go.Response[v1.CreateStudentResponse], err error) {
	if req.Msg.FullName != "" {
		return &connect_go.Response[v1.CreateStudentResponse]{
			Msg: &v1.CreateStudentResponse{
				Id:       "12321",
				FullName: req.Msg.FullName,
			},
		}, nil
	}
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("student.v1.StudentService.CreateStudent is not implemented"))
}
