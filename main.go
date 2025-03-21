package main

import (
	"fmt"
	"golang-dataon/configs"
	"golang-dataon/databases"
	"golang-dataon/middlewares"
	"golang-dataon/repository"
	"golang-dataon/server"
	"golang-dataon/service"
	"golang-dataon/specs"
	"golang-dataon/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	configs := configs.LoadConfig()
	repository := repository.NewRepository(databases.NewGormDB(configs))
	service := service.NewService(repository, configs)
	var server specs.ServerInterface = server.NewServer(configs, service)
	e := echo.New()

	specs.RegisterHandlers(e, server)
	e.Validator = utils.NewEchoValidator()
	e.Use(middlewares.ErrorHandler)

	e.Static("/", "statics")
	e.GET("/form", service.RenderAddForm)
	e.GET("/form/:id", service.RenderEditForm)
	e.GET("/form/:id/delete", service.RenderDeleteForm)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", configs.PORT)))
}
