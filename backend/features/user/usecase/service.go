package usecase

import (
	"institute/features/user"
	"institute/features/user/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model user.Repository
}

func New(model user.Repository) user.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResUser {
	var users []dtos.ResUser

	usersEnt := svc.model.Paginate(page, size)

	for _, user := range usersEnt {
		var data dtos.ResUser

		if err := smapping.FillStruct(&data, smapping.MapFields(user)); err != nil {
			log.Error(err.Error())
		} 
		
		users = append(users, data)
	}

	return users
}

func (svc *service) FindByID(userID int) *dtos.ResUser {
	res := dtos.ResUser{}
	user := svc.model.SelectByID(userID)

	if user == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(user))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newUser dtos.InputUser) *dtos.ResUser {
	user := user.User{}
	
	err := smapping.FillStruct(&user, smapping.MapFields(newUser))
	if err != nil {
		log.Error(err)
		return nil
	}

	userID := svc.model.Insert(user)

	if userID == -1 {
		return nil
	}

	resUser := dtos.ResUser{}
	errRes := smapping.FillStruct(&resUser, smapping.MapFields(newUser))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resUser
}

func (svc *service) Modify(userData dtos.InputUser, userID int) bool {
	newUser := user.User{}

	err := smapping.FillStruct(&newUser, smapping.MapFields(userData))
	if err != nil {
		log.Error(err)
		return false
	}

	newUser.ID = userID
	rowsAffected := svc.model.Update(newUser)

	if rowsAffected <= 0 {
		log.Error("There is No User Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(userID int) bool {
	rowsAffected := svc.model.DeleteByID(userID)

	if rowsAffected <= 0 {
		log.Error("There is No User Deleted!")
		return false
	}

	return true
}