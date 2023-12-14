package course

import (
	"institute/features/course/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Course
	Insert(newCourse Course) int64
	SelectByID(courseID int) *Course
	Update(course Course) int64
	DeleteByID(courseID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResCourse
	FindByID(courseID int) *dtos.ResCourse
	Create(newCourse dtos.InputCourse) *dtos.ResCourse
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
