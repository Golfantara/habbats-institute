package handler

import (
	"institute/helpers"
	helper "institute/helpers"
	"strconv"

	"institute/features/user"
	"institute/features/user/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.Usecase
}

func New(service user.Usecase) user.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetUsers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		users := ctl.service.FindAll(page, size)

		if users == nil {
			return ctx.JSON(404, helper.Response("There is No Users!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": users,
		}))
	}
}


func (ctl *controller) UserDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		userID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helper.Response("User Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": user,
		}))
	}
}

func (ctl *controller) CreateUser() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputUser{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		user := ctl.service.Create(input)

		if user == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": user,
		}))
	}
}

func (ctl *controller) UpdateUser() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputUser{}

		userID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helper.Response("User Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, userID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("User Success Updated!"))
	}
}

func (ctl *controller) DeleteUser() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		userID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		user := ctl.service.FindByID(userID)

		if user == nil {
			return ctx.JSON(404, helper.Response("User Not Found!"))
		}

		delete := ctl.service.Remove(userID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("User Success Deleted!", nil))
	}
}
