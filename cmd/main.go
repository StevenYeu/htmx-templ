package main

import (
	"github.com/StevenYeu/fem-htmx/models"
	"github.com/StevenYeu/fem-htmx/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	page := models.NewPage()
	e.GET("/", func(c echo.Context) error {
		return Render(c, 200, views.Index(page))
	})
	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.HasEmail(email) {
			formData := models.NewFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"
			return Render(c, 422, views.Form(formData))
		}

		contact := models.NewContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)
        Render(c,200,views.Form(models.NewFormData()))
        return Render(c, 200, views.OOBContact(contact))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
