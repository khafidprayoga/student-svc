package svc

import (
	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
	"go.uber.org/zap"
)

type StudentServiceServerImpl struct {
	Log *zap.Logger
	studentv2.UnimplementedStudentServiceServer
}
