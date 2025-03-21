package server

import (
	"golang-dataon/specs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Add(c echo.Context) (err error) {
	defer c.Request().Body.Close()

	var reqBody *specs.AddJSONRequestBody
	if reqBody, err = getRequestBody[specs.AddJSONRequestBody](c); err != nil {
		return
	}

	result, err := s.Service.Add(reqBody)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (s *Server) GetAll(c echo.Context) (err error) {
	result, err := s.Service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) GetDetails(c echo.Context, id specs.IdPath) (err error) {
	result, err := s.Service.GetDetails(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) Edit(c echo.Context, id specs.IdPath) (err error) {
	defer c.Request().Body.Close()

	var reqBody *specs.EditJSONRequestBody
	if reqBody, err = getRequestBody[specs.EditJSONRequestBody](c); err != nil {
		return
	}
	result, err := s.Service.Edit(id, reqBody)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (s *Server) Delete(c echo.Context, id specs.IdPath) (err error) {
	err = s.Service.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Deleted"})
}
