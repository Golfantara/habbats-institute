package repository

import (
	"institute/features/course"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) course.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []course.Course {
	var courses []course.Course

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&courses)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return courses
}

func (mdl *model) Insert(newCourse course.Course) int64 {
	result := mdl.db.Create(&newCourse)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newCourse.ID)
}

func (mdl *model) SelectByID(courseID int) *course.Course {
	var course course.Course
	result := mdl.db.First(&course, courseID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &course
}

func (mdl *model) Update(course course.Course) int64 {
	result := mdl.db.Save(&course)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(courseID int) int64 {
	result := mdl.db.Delete(&course.Course{}, courseID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}