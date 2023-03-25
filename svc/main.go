package svc

import (
	"github.com/khafidprayoga/student-svc/common/data"
	studentv2 "github.com/khafidprayoga/student-svc/gen/student/v2"
	"go.uber.org/zap"
)

type StudentServiceServerImpl struct {
	Log  *zap.Logger
	Data *data.StudentDataRepository
	studentv2.UnimplementedStudentServiceServer
}
