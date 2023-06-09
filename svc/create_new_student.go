package svc

import (
	"context"
	"regexp"

	"buf.build/gen/go/khafidprayoga/student-svc/protocolbuffers/go/student/v2"
	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/khafidprayoga/student-svc/common/constant"
	"github.com/khafidprayoga/student-svc/common/model"

	"go.uber.org/zap"
)

func (server *StudentServiceServerImpl) CreateStudent(
	ctx context.Context, req *studentv2.CreateStudentRequest,
) (res *studentv2.CreateStudentResponse, err error) {
	errValidate := validation.ValidateStruct(
		req,
		validation.Field(&req.FullName, validation.Required,
			validation.Length(5, 32),
		),
		validation.Field(&req.Email, validation.Required, validation.Match(
			regexp.MustCompile(constant.EmailPattern),
		)),
		validation.Field(&req.Address, validation.Required),
	)

	if errValidate != nil {
		return nil, errValidate
	}

	dob := req.BirthDate.AsTime()
	newStudent := &model.Student{
		FullName:    req.FullName,
		Address:     req.Address,
		Hobbies:     req.Hobbies,
		BirthDate:   &dob,
		Gender:      studentv2.GenderType_value[req.Gender.String()],
		Email:       req.Email,
		Nationality: studentv2.StudentNationality_value[req.Nationality.String()],
	}

	studentId, errInsert := server.Data.CreateStudent(newStudent)
	if err != nil {
		return nil, errInsert
	}

	server.Log.Info("creating new student",
		zap.String("full_name", req.GetFullName()),
		zap.String("id", studentId),
	)

	return &studentv2.CreateStudentResponse{
		Id:          studentId,
		FullName:    req.FullName,
		BirthDate:   req.GetBirthDate(),
		Gender:      studentv2.GenderType_value[req.GetGender().String()],
		Address:     req.GetAddress(),
		Hobbies:     req.GetHobbies(),
		Nationality: studentv2.StudentNationality_value[req.GetNationality().String()],
		Email:       req.GetEmail(),
	}, nil
}
