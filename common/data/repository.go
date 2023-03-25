package data

import (
	"errors"
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/google/uuid"
	"github.com/khafidprayoga/student-svc/common/model"
	"github.com/khafidprayoga/student-svc/gen/student/v2"
)

var studentData = make(map[string]model.Student)

type StudentDataRepository struct{}

func New() *StudentDataRepository {
	return &StudentDataRepository{}
}

func init() {
	now := time.Now()

	WIB, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Printf("failed to init seed data %v", err)
	}
	dob := time.Date(1996, time.June, 9, 0, 0, 0, 0, time.UTC)
	dob = dob.In(WIB)

	// Seed initial data
	student := model.Student{
		Id:          "seed",
		FullName:    "Jane Doe",
		Address:     "Jalan Jend. Sudirman No. 321",
		Hobbies:     []string{"listening to the music", "singing"},
		BirthDate:   &dob,
		Gender:      studentv2.GenderType_value[studentv2.GenderType_GENDER_TYPE_FEMALE.String()],
		Email:       "janedoe@example.com",
		Nationality: studentv2.StudentNationality_value[studentv2.StudentNationality_STUDENT_NATIONALITY_FOREIGN.String()],
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	studentData[student.Id] = student
}

func (data *StudentDataRepository) CreateStudent(s *model.Student) (id string, err error) {
	now := time.Now()
	preparedData := model.Student{
		Id:          uuid.NewString(),
		FullName:    s.FullName,
		Address:     s.Address,
		Hobbies:     s.Hobbies,
		BirthDate:   s.BirthDate,
		Gender:      s.Gender,
		Email:       s.Email,
		Nationality: s.Nationality,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	studentData[preparedData.Id] = preparedData
	if studentData[preparedData.Id].Id == "" {
		return "", fmt.Errorf("failed to insert student %v ", s.FullName)
	}
	return preparedData.Id, nil
}

func (data *StudentDataRepository) GetDetailStudent(
	studentId string,
) (*model.Student, error) {
	details := studentData[studentId]
	if details.Id == "" {
		return nil, errors.New("student not found")
	}
	return &details, nil
}

func (data *StudentDataRepository) DeleteStudent(studentId string) error {
	delete(studentData, studentId)
	if studentData[studentId].Id != "" {
		return fmt.Errorf(
			"failed to delete student wit id %v", studentId,
		)
	}

	return nil
}

func (data *StudentDataRepository) GetListStudent() (
	listStudent []model.Student,
) {
	for _, student := range studentData {
		listStudent = append(listStudent, student)
	}
	return
}
