package course

import (
	"institute/features/course/dtos"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Course
	Insert(newCourse *Course) (*Course, error)
	SelectByID(courseID int) *Course
	Update(course Course) int64
	DeleteByID(courseID int) int64
	UploadFile(fileHeader *multipart.FileHeader, name string) (string, error)
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResCourse
	FindByID(courseID int) *dtos.ResCourse
	Create(newCourse dtos.InputCourse, file *multipart.FileHeader) (*dtos.ResCourse, error)
	Modify(courseData dtos.InputCourse, courseID int) bool
	Remove(courseID int) bool
}

type Handler interface {
	GetCourses() echo.HandlerFunc
	CourseDetails() echo.HandlerFunc
	CreateCourse() echo.HandlerFunc
	UpdateCourse() echo.HandlerFunc
	DeleteCourse() echo.HandlerFunc
}
