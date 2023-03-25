package svc

import (
	"context"
	"errors"

	"github.com/khafidprayoga/student-svc/gen/student/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *StudentServiceServerImpl) GetDetailStudent(
	ctx context.Context, req *studentv2.GetDetailStudentRequest,
) (
	res *studentv2.GetDetailStudentResponse, err error) {
	studentId := req.StudentId
	if studentId == "" {
		return nil, errors.New("student id must be not empty string")
	}

	data, err := server.Data.GetDetailStudent(studentId)
	if err != nil {
		return nil, err
	}

	res = &studentv2.GetDetailStudentResponse{
		Id:          data.Id,
		FullName:    data.FullName,
		BirthDate:   timestamppb.New(*data.BirthDate),
		Gender:      data.Gender,
		Address:     data.Address,
		Hobbies:     data.Hobbies,
		Nationality: data.Nationality,
		Email:       data.Email,
		CreatedAt:   timestamppb.New(*data.CreatedAt),
		UpdatedAt:   timestamppb.New(*data.UpdatedAt),
	}
	return res, nil
}
