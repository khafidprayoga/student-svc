package svc

import (
	"context"
	"errors"
	"fmt"

	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
)

func (server *StudentServiceServerImpl) DeleteStudent(
	ctx context.Context, req *studentv2.DeleteStudentRequest) (
	res *studentv2.DeleteStudentResponse, err error,
) {
	studentId := req.GetStudentId()
	if studentId == "" {
		err = errors.New("student id must be not empty string")
		return
	}

	_, errGetDetail := server.Data.GetDetailStudent(studentId)
	if errGetDetail != nil {
		err = errGetDetail
		return
	}
	errDeleteStudent := server.Data.DeleteStudent(studentId)
	if errDeleteStudent != nil {
		err = errDeleteStudent
		return
	}

	return &studentv2.DeleteStudentResponse{
		Message: fmt.Sprintf(
			"success deleted student with id %v", studentId,
		),
	}, nil
}
