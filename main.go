package main

import (
	"log"

	"crudgengui/controller"
	"crudgengui/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
  var err error
  
	e := echo.New()
	e.Use(middleware.Static("/static"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}:${status}, uri:\"${uri}\", path:\"${path}\", error:\"${error}\"\n",
	}))

  // Create templates
  templates := controller.NewTemplateRegistry()

  // base template
  err=templates.AddTemplate("base.html","","template/base/side_navigation.html", "template/base/delete_popup.html", "template/base/base.html",  "template/base/side_navigation.html", "template/base/top_navigation.html")
  if err!=nil {
    log.Panic(err)
  }  
  
  err=templates.AddTemplate("entities.html","base.html", "template/entity_popup.html", "template/base/script.html", "template/entities.html")
  if err!=nil {
    log.Panic(err)
  }  

  err=templates.AddTemplate("lookups.html","base.html", "template/lookup_popup.html", "template/base/script.html", "template/lookups.html")
  if err!=nil {
    log.Panic(err)
  }  

  err=templates.AddTemplate("lookup.html","base.html", "template/lookupadd_popup.html", "template/base/script.html", "template/lookup.html")
  if err!=nil {
    log.Panic(err)
  }  
  
  err=templates.AddTemplate("relations.html","base.html", "template/relation_popup.html", "template/base/script.html", "template/relations.html")
  if err!=nil {
    log.Panic(err)
  }  
  
  err=templates.AddTemplate("entity.html","base.html","template/field_popup.html", "template/base/script.html", "template/entity.html")
  if err!=nil {
    log.Panic(err)
  }  

  err=templates.AddTemplate("field.html","base.html","template/field.html")
  if err!=nil {
    log.Panic(err)
  }  
  
  err=templates.AddTemplate("index.html","base.html")
  if err!=nil {
    log.Panic(err)
  }  
	e.Renderer = templates

  // Create Repository
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

	e.GET("/fields/:id", mc.ShowField) // :id is the entityname where the field belongs 
  e.POST("/fields", mc.InsertField)
	e.POST("/fields/:id", mc.DeleteField) // :id is the entityname where the field belongs to

	e.GET("/lookups", mc.ShowAllLookups)
	e.POST("/lookups", mc.InsertLookup)
  e.GET("/lookups/:id", mc.ShowLookup)
	e.POST("/lookups/:id", mc.ModifyLookup) // save, delete Lookup

  
	e.Logger.Fatal(e.Start(":1323"))
}

