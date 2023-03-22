package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"

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

func main() {
	e := echo.New()
	e.Use(middleware.Static("/static"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}, path=${path}\n",
	}))

	err := controller.LoadModel("model.yaml")
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"title":     strings.Title,
		"lowercase": strings.ToLower,
		"uppercase": strings.ToUpper,
	}

	templates := make(map[string]*template.Template)

	// base template
	templates["base.html"] = template.Must(template.New("base.html").Funcs(funcMap).ParseFiles("template/base/side_navigation.html", "template/base/delete_popup.html", "template/base/base.html", "template/entity_popup.html", "template/relation_popup.html", "template/base/side_navigation.html", "template/base/top_navigation.html"))

	templates["entities.html"], err = template.Must(templates["base.html"].Clone()).ParseFiles("template/entities.html")
	if err != nil {
		log.Fatal(err)
	}
	templates["relations.html"], err = template.Must(templates["base.html"].Clone()).ParseFiles("template/relations.html")
	if err != nil {
		log.Fatal(err)
	}
	templates["index.html"] = template.Must(templates["base.html"].Clone())

	templates["entity.html"], err = template.Must(templates["base.html"].Clone()).ParseFiles("template/field_popup.html", "template/entity.html")
	if err != nil {
		log.Fatal(err)
	}

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", showDashboard)
	e.GET("/entities/:id", showEntity)
	e.POST("/entities/:id", deleteEntity)
	e.GET("/entities", showAllEntities)
	e.POST("/entities", insertEntity)

	e.GET("/relations/:id", showRelation)
	e.POST("/relations/:id", deleteRelation)
	e.GET("/relations", showAllRelations)
	e.POST("/relations", insertRelation)

	e.POST("/fields", insertField)
	e.POST("/fields/:id", deleteField)
	e.Logger.Fatal(e.Start(":1323"))
}

func showDashboard(c echo.Context) error {
	m := controller.GetModel()
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"model": m,
		"title": "Home",
	})
}

// ------------- Entities -------------

// showAllEntities
func showAllEntities(c echo.Context) error {
	m := controller.GetModel()
	return c.Render(http.StatusOK, "entities.html", map[string]interface{}{
		"model": m,
		"title": "Entities",
	})
}

// insertEntity
func insertEntity(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	e := model.NewEntity()
	if err := c.Bind(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if (len(e.Name) < 2) || (len(e.Type) < 2) {
		return nil
	}
	if _, ok := controller.GetEntity(e.Name); ok {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Name %v is already in use", e.Name))
	}

	if err := controller.SaveOrUpdateEntity(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return showAllEntities(c)
}

// showEntity shows detail page to model or edit screen for new model
func showEntity(c echo.Context) error {

	id := c.Param("id")
	entity, ok := controller.GetEntity(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	all := controller.GetAllEntities()
	return c.Render(http.StatusOK, "entity.html", map[string]interface{}{
		"model":  all,
		"entity": entity,
		"title":  fmt.Sprint("Entity: ", entity.Name),
	})
}

// deleteEntity shows detail page to model or edit screen for new model
func deleteEntity(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if _, ok := controller.GetEntity(id); !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	if err := controller.DeleteEntity(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/entities")
}

// ------------- Relations -------------

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

	if _, ok := controller.GetRelation(id); !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	if err := controller.DeleteRelation(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/relations")

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

	if len(r.Name) < 3 {
		return nil
	}

	if _, ok := controller.GetEntity(r.Name); ok {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Name %v is already in use", r.Name))
	}

	if err := controller.SaveOrUpdateRelation(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return showAllEntities(c)
}

// ------------- Fields -------------

// deleteRelation shows detail page to model or edit screen for new model
func deleteField(c echo.Context) error {
	// User ID from path `users/:id`
	fname := c.FormValue("field_name")
	ename := c.FormValue("entity_name")

	if (len(fname) < 2) || (len(ename) < 2) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity '%s', field '%s' not found", ename, fname))
	}

	if err := controller.DeleteField(ename, fname); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return showAllEntities(c)

}

// insertRelation
func insertField(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	ename := c.FormValue("entity_name")

	field := new(model.Field)
	if err := c.Bind(field); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if (len(field.Name) < 3) || (len(ename) < 3) {
		return nil
	}

	if err := controller.SaveOrUpdateField(ename,field); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return showAllEntities(c)
}
