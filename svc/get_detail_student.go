package svc

import (
	"context"

	"github.com/google/uuid"
	"github.com/khafidprayoga/student-svc/gen/student/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *StudentServiceServerImpl) GetDetailStudent(
	ctx context.Context, req *studentv2.GetDetailStudentRequest,
) (

	res *studentv2.GetDetailStudentResponse, err error) {
	res = &studentv2.GetDetailStudentResponse{
		Students: &studentv2.Student{
			Id:          uuid.NewString(),
			FullName:    "Akbar Maulana",
			BirthDate:   timestamppb.Now(),
			Gender:      studentv2.GenderType_GENDER_TYPE_MALE,
			Address:     "Jl. Soekarno Hatta",
			Hobbies:     []string{"Fishing", "Listening to the music"},
			Nationality: studentv2.StudentNationality_STUDENT_NATIONALITY_CITIZEN,
		},
	}
	return res, nil
}
