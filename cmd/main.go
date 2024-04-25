package main

import (
	"html/template"
	"io"
	"strconv"

	"github.com/StevenYeu/fem-htmx/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{templates: template.Must(template.ParseGlob("views/*.html"))}
}

type Count struct {
	Count int
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	count := Count{Count: 0}
	e.GET("/", func(c echo.Context) error {
		parsedCount := strconv.Itoa(count.Count)
		return Render(c, 200, views.Index(parsedCount))
	})
	e.POST("/count", func(c echo.Context) error {
		count.Count++
		parsedCount := strconv.Itoa(count.Count)
		return Render(c, 200, views.Count(parsedCount))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
