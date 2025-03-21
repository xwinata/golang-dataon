package server

import (
	"golang-dataon/configs"
	"golang-dataon/customErrors"
	"golang-dataon/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Configs *configs.Config
	Service service.IService
}

func NewServer(configs *configs.Config, service service.IService) *Server {
	return &Server{
		Configs: configs,
		Service: service,
	}
}

// handler extention to bind and validate request body payload.
func getRequestBody[T interface{}](c echo.Context) (reqBody *T, err error) {
	reqBody = new(T)
	if err = c.Bind(&reqBody); err != nil {
		return reqBody, &customErrors.ApiError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	if err = c.Validate(reqBody); err != nil {
		return reqBody, &customErrors.ApiError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	return
}
