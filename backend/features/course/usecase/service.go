package usecase

import (
	"institute/features/course"
	"institute/features/course/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model course.Repository
}

func New(model course.Repository) course.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResCourse {
	var courses []dtos.ResCourse

	coursesEnt := svc.model.Paginate(page, size)

	for _, course := range coursesEnt {
		var data dtos.ResCourse

		if err := smapping.FillStruct(&data, smapping.MapFields(course)); err != nil {
			log.Error(err.Error())
		} 
		
		courses = append(courses, data)
	}

	return courses
}

func (svc *service) FindByID(courseID int) *dtos.ResCourse {
	res := dtos.ResCourse{}
	course := svc.model.SelectByID(courseID)

	if course == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(course))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newCourse dtos.InputCourse) *dtos.ResCourse {
	course := course.Course{}
	
	err := smapping.FillStruct(&course, smapping.MapFields(newCourse))
	if err != nil {
		log.Error(err)
		return nil
	}

	courseID := svc.model.Insert(course)

	if courseID == -1 {
		return nil
	}

	resCourse := dtos.ResCourse{}
	errRes := smapping.FillStruct(&resCourse, smapping.MapFields(newCourse))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resCourse
}

func (svc *service) Modify(courseData dtos.InputCourse, courseID int) bool {
	newCourse := course.Course{}

	err := smapping.FillStruct(&newCourse, smapping.MapFields(courseData))
	if err != nil {
		log.Error(err)
		return false
	}

	newCourse.ID = courseID
	rowsAffected := svc.model.Update(newCourse)

	if rowsAffected <= 0 {
		log.Error("There is No Course Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(courseID int) bool {
	rowsAffected := svc.model.DeleteByID(courseID)

	if rowsAffected <= 0 {
		log.Error("There is No Course Deleted!")
		return false
	}

	return true
}