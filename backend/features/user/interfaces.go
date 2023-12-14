package user

import (
	"institute/features/user/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []User
	Insert(newUser User) int64
	SelectByID(userID int) *User
	Update(user User) int64
	DeleteByID(userID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResUser
	FindByID(userID int) *dtos.ResUser
	Create(newUser dtos.InputUser) *dtos.ResUser
	Modify(userData dtos.InputUser, userID int) bool
	Remove(userID int) bool
}

type Handler interface {
	GetUsers() echo.HandlerFunc
	UserDetails() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}
