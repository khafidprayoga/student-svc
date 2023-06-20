package svc

import (
	"buf.build/gen/go/khafidprayoga/student-svc/grpc/go/student/v2/studentv2grpc"

	"github.com/khafidprayoga/student-svc/common/data"

	"go.uber.org/zap"
)

type StudentServiceServerImpl struct {
	Log  *zap.Logger
	Data *data.StudentDataRepository
	studentv2grpc.UnimplementedStudentServiceServer
}
