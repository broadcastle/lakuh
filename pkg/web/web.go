package web

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"broadcastle.co/code/lakuh/pkg/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var e *echo.Echo

// Web starts the web interface.
func Web() {

	db.Init()
	defer db.Close()

	e = echo.New()

	// e.HideBanner = true
	e.Static("/", "public")
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Renderer = &Template{templates: template.Must(template.ParseGlob("resources/templates/*"))}

	// API

	api1()

	port := viper.GetInt("lakuh.port")

	logrus.Fatal(e.Start(":" + strconv.Itoa(port)))

}

func empty(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "empty path")
}

func useJWT(c *echo.Group) {
	c.Use(middleware.JWT([]byte(viper.GetString("lakuh.token"))))
}

// Template for the renderer
type Template struct {
	templates *template.Template
}

// Render the template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
