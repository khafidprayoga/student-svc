package svc

import (
	"context"

	"github.com/khafidprayoga/student-svc/gen/student/v2"
)

func (server *StudentServiceServerImpl) CreateStudent(
	ctx context.Context, req *studentv2.CreateStudentRequest,
) (res *studentv2.CreateStudentResponse, err error) {
	return nil, err
}
