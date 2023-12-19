package main

import (
	"fmt"
	"institute/config"
	"institute/features/auth"
	"institute/features/course"
	"institute/helpers"
	"institute/middlewares"
	"institute/routes"
	"institute/utils"

	"github.com/labstack/echo/v4"

	ah "institute/features/auth/handler"
	ar "institute/features/auth/repository"
	au "institute/features/auth/usecase"

	ch "institute/features/course/handler"
	cr "institute/features/course/repository"
	cu "institute/features/course/usecase"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	jwtService := helpers.NewJWT(*cfg)
	
	middlewares.LogMiddlewares(e)
	routes.Auth(e, AuthHandler(), jwtService, *cfg)
	routes.Courses(e, CourseHandler())

	e.Start(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func AuthHandler() auth.Handler{
	config := config.InitConfig()

	db := utils.InitDB()
	jwt := helpers.NewJWT(*config)
	hash := helpers.NewHash()
	validation := helpers.NewValidationRequest()

	repo := ar.New(db)
	uc := au.New(repo, jwt, hash, validation)
	return ah.New(uc)
}

func CourseHandler() course.Handler {
	config := config.InitConfig()
	cdn := utils.CloudinaryInstance(*config)

	db := utils.InitDB()

	repo := cr.New(db, cdn, config)
	uc :=	cu.New(repo)
	return ch.New(uc)
}