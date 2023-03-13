package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	model "crudgengui/model"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var m *model.Model

func main() {
	e := echo.New()
	e.Use(middleware.Static("/static"))

	m = model.NewModel()
	file, err := os.Open("model.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = m.ReadYAML(file)
	if err != nil {
		log.Fatal(err)
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", showDashboard)
	e.GET("/models/:id", getSingleModel)
	e.GET("/models", getModels)
	e.POST("/entities", insertEntity)
	e.POST("/relations", insertRelation)
	e.Logger.Fatal(e.Start(":1323"))
}

func showDashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
	//	return c.File("public/index.html")
}

// getSingleModel shows detail page to model or edit screen for new model
func getSingleModel(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	if id != "0" {
		return c.String(http.StatusOK, id)
	} else {
		return c.String(http.StatusOK, "new")
	}
}

// getSingleModel
func getModels(c echo.Context) error {
	return c.String(http.StatusOK, "List all models")
}

// insertEntity
func insertEntity(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	m := new(model.Entity)
	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, showDashboard(c))
}

// insertRelation
func insertRelation(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	r := new(model.Relation)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, showDashboard(c))
}
