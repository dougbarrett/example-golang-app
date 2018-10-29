package server

import (
	"html/template"
	"io"

	"github.com/dougbarrett/example-golang-app/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type server struct {
	handler handler.Service
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewFrontend(
	handler handler.Service,
) (
	*echo.Echo,
	error,
) {
	s := &server{}
	s.handler = handler

	mux := echo.New()

	// Middleware
	mux.Use(middleware.Logger())
	mux.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("server/views/*.html")),
	}
	mux.Renderer = t

	mux.GET("/", s.getAllEntries)
	mux.GET("/entries/new", s.newEntry)
	mux.POST("/entries/create", s.createEntry)
	mux.POST("/entries/delete", s.deleteEntry)

	return mux, nil
}
