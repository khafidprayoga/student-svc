package svc

import (
	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
)

type StudentServiceServerImpl struct {
	studentv2.UnimplementedStudentServiceServer
}
