package routes

import (
	"institute/features/course"

	"github.com/labstack/echo/v4"
)

func Courses(e *echo.Echo, handler course.Handler) {
	courses := e.Group("/courses")

	courses.GET("", handler.GetCourses())
	courses.POST("", handler.CreateCourse())
	
	courses.GET("/:id", handler.CourseDetails())
	courses.PUT("/:id", handler.UpdateCourse())
	courses.DELETE("/:id", handler.DeleteCourse())
}