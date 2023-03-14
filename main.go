package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
  "strings"

	"crudgengui/controller"
	model "crudgengui/model"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return fmt.Errorf("Template not found -> %s", name)
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

/*
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
*/

func main() {
	e := echo.New()
	e.Use(middleware.Static("/static"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	err := controller.LoadModel("model.yaml")
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"title": strings.Title,
    "lowercase": strings.ToLower,
    "uppercase": strings.ToUpper,
	}
  
	templates := make(map[string]*template.Template)

  // base template
	templates["base.html"] = template.Must(template.New("base.html").Funcs(funcMap).ParseFiles("template/base/side_navigation.html", "template/base/base.html", "template/entity_popup.html", "template/relation_popup.html", "template/base/side_navigation.html", "template/base/top_navigation.html"))

	templates["entities.html"], err = template.Must(templates["base.html"].Clone()).ParseFiles("template/entities.html")
 	if err != nil {
		log.Fatal(err)
	}    
  templates["relations.html"], err = template.Must(templates["base.html"].Clone()).ParseFiles("template/relations.html")
 	if err != nil {
		log.Fatal(err)
	}    
  templates["index.html"] = template.Must(templates["base.html"].Clone())                                    

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", showDashboard)
	e.GET("/entities/:id", showEntity)
	e.DELETE("/entities/:id", deleteEntity)
	e.GET("/entities", showAllEntities)
	e.POST("/entities", insertEntity)

	e.GET("/relations/:id", showRelation)
	e.DELETE("/relations/:id", deleteRelation)
	e.GET("/relations", showAllRelations)
	e.POST("/relations", insertRelation)
	e.Logger.Fatal(e.Start(":1323"))
}

func showDashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
	//	return c.File("public/index.html")
}

// showAllEntities
func showAllEntities(c echo.Context) error {
	m := controller.GetModel()
	return c.Render(http.StatusOK, "entities.html", map[string]interface{}{
		"model": m,
		"title": "Entities",
	})
}

// showEntity shows detail page to model or edit screen for new model
func showEntity(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// deleteEntity shows detail page to model or edit screen for new model
func deleteEntity(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if err := controller.DeleteEntity(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusAccepted, showDashboard(c))
}

// showAllRelations
func showAllRelations(c echo.Context) error {
	m := controller.GetModel()
	return c.Render(http.StatusOK, "relations.html", map[string]interface{}{
		"model": m,
		"title": "Relations",
	})
}

// showRelation shows detail page to model or edit screen for new model
func showRelation(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// deleteRelation shows detail page to model or edit screen for new model
func deleteRelation(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if err := controller.DeleteEntity(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusAccepted, showDashboard(c))

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
	if err := controller.SaveOrUpdateEntity(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, showAllEntities(c))
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
	r.Name = r.Source + r.Type + r.Destination
	if err := controller.SaveOrUpdateRelation(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, showDashboard(c))
}
