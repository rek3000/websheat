package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
  "websheat/internal/handlers"
  "html/template"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	t := &Template{
		templates: template.Must(template.ParseGlob("./internal/templates/**/*.html")),
	}
	e.Renderer = t

  e.GET("/", handlers.HandleHome)
  e.GET("/todos", handlers.HandleGetTodos)
  e.POST("/todos", handlers.HandleCreateTodo)
  e.DELETE("/todos/:id", handlers.HandleDeleteTodo)
  e.PUT("/todos/:id/toggle", handlers.HandleToggleTodo)
	e.GET("/great", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Rek")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
