package svc

import (
	"context"

	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *StudentServiceServerImpl) GetListStudent(
	ctx context.Context, req *studentv2.GetListStudentRequest,
) (res *studentv2.GetListStudentResponse, err error, ) {
	var (
		finalData []*studentv2.GetDetailStudentResponse
	)

	studentList := server.Data.GetListStudent()
	if studentList != nil {
		for _, each := range studentList {
			format := studentv2.GetDetailStudentResponse{
				Id:          each.Id,
				FullName:    each.FullName,
				BirthDate:   timestamppb.New(*each.BirthDate),
				Gender:      each.Gender,
				Address:     each.Address,
				Hobbies:     each.Hobbies,
				Nationality: each.Nationality,
				Email:       each.Email,
				CreatedAt:   timestamppb.New(*each.CreatedAt),
				UpdatedAt:   timestamppb.New(*each.UpdatedAt),
			}

			finalData = append(finalData, &format)
		}
	}

	return &studentv2.GetListStudentResponse{
		Details: finalData,
	}, nil

}
