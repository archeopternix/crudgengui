package main

import (
	"html/template"
	"log"

	"crudgengui/controller"
	"crudgengui/repository"

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
  var err error
  
	e := echo.New()
	e.Use(middleware.Static("/static"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "HTTP/${method}: ${status}, uri=${uri}, error=${error}, path=${path}\n",
	}))

  /*
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
 */
  templates := controller.NewTemplateRegistry()

// base template
err=templates.AddTemplate("base.html","","template/base/side_navigation.html", "template/base/delete_popup.html", "template/base/base.html", "template/entity_popup.html", "template/relation_popup.html", "template/base/side_navigation.html", "template/base/top_navigation.html")
if err!=nil {
  log.Panic(err)
}  
  
err=templates.AddTemplate("entities.html","base.html","template/entities.html")
 templates.AddTemplate("relations.html","base.html","template/relations.html")
if err!=nil {
  log.Panic(err)
}  
  err=templates.AddTemplate("entity.html","base.html","template/field_popup.html","template/entity.html")
if err!=nil {
  log.Panic(err)
}  

    err=templates.AddTemplate("index.html","base.html")
if err!=nil {
  log.Panic(err)
}  
  
	e.Renderer = templates

  mc := controller.NewModelController(repository.NewModelRepository(repository.NewYAMLModel("model.yaml")))

	e.GET("/", mc.ShowDashboard)
	e.GET("/entities/:id", mc.ShowEntity)
	e.POST("/entities/:id", mc.DeleteEntity)
	e.GET("/entities", mc.ShowAllEntities)
	e.POST("/entities", mc.InsertEntity)

	e.GET("/relations/:id", mc.ShowRelation)
	e.POST("/relations/:id", mc.DeleteRelation)
	e.GET("/relations", mc.ShowAllRelations)
	e.POST("/relations", mc.InsertRelation)

	e.POST("/fields", mc.InsertField)
	e.POST("/fields/:id", mc.DeleteField)
	e.Logger.Fatal(e.Start(":1323"))
}

