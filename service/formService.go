package service

import (
	"text/template"

	"github.com/labstack/echo/v4"
)

type FormData struct {
	Value  string
	Upline string
	Error  string
}

func (s *Service) RenderAddForm(ctx echo.Context) error {
	tmpl, err := template.ParseFiles("templates/addform.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(ctx.Response(), FormData{})
}

func (s *Service) RenderEditForm(ctx echo.Context) error {
	tmpl, err := template.ParseFiles("templates/editform.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(ctx.Response(), FormData{})
}

func (s *Service) RenderDeleteForm(ctx echo.Context) error {
	tmpl, err := template.ParseFiles("templates/deleteform.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(ctx.Response(), FormData{})
}
