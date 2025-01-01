package internal

import (
	controller "crudgengui/internal/controller"
	repository "crudgengui/internal/repository"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	base = "../../"
	tpl  = "../../internal/template/"
)

type GuiServer struct {
	e *echo.Echo
}

func NewGuiServer() GuiServer {
	s := GuiServer{e: echo.New()}
	s.e.Use(middleware.Static("/static"))
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}:${status}, uri:\"${uri}\", path:\"${path}\", error:\"${error}\"\n",
	}))
	return s
}

func (s *GuiServer) Init() {
	s.setTemplates()
	s.setRoutes()
}

func (s *GuiServer) setTemplates() {
	// Create templates
	templates := NewTemplateRegistry()

	// Define templates with their files
	templateDefinitions := []struct {
		name  string
		base  string
		files []string
	}{
		{"base.html", "", []string{tpl + "base/side_navigation.html", tpl + "base/delete_popup.html", tpl + "base/base.html", tpl + "base/side_navigation.html", tpl + "base/top_navigation.html"}},
		{"entities.html", "base.html", []string{tpl + "entity_popup.html", tpl + "base/script.html", tpl + "entities.html"}},
		{"lookups.html", "base.html", []string{tpl + "lookup_popup.html", tpl + "base/script.html", tpl + "lookups.html"}},
		{"lookup.html", "base.html", []string{tpl + "lookupadd_popup.html", tpl + "base/script.html", tpl + "lookup.html"}},
		{"relations.html", "base.html", []string{tpl + "relation_popup.html", tpl + "base/script.html", tpl + "relations.html"}},
		{"entity.html", "base.html", []string{tpl + "base/script.html", tpl + "entity.html"}},
		{"field.html", "base.html", []string{tpl + "field.html"}},
		{"index.html", "base.html", []string{tpl + "index.html"}},
		{"project.html", "base.html", []string{tpl + "project.html"}},
	}

	// Add templates using the helper function
	for _, tmpl := range templateDefinitions {
		templates.AddTemplateOrPanic(tmpl.name, tmpl.base, tmpl.files...)
	}

	s.e.Renderer = templates
}

func (s *GuiServer) setRoutes() {
	// Create Repository
	mc := controller.NewModelController(repository.NewModelRepository(repository.NewYAMLModel(base + "data/model.yaml")))

	s.e.GET("/", mc.ShowDashboard)
	s.e.GET("/project", mc.ShowProject)
	s.e.POST("/project", mc.SaveProject)

	// Group for "/entities" routes
	entitiesGroup := s.e.Group("/entities")
	{
		entitiesGroup.GET("/:id", mc.ShowEntity)
		entitiesGroup.POST("/:id", mc.DeleteEntity)
		entitiesGroup.GET("", mc.ShowAllEntities)
		entitiesGroup.POST("", mc.InsertEntity)
	}

	// Group for "/relations" routes
	relationsGroup := s.e.Group("/relations")
	{
		relationsGroup.GET("/:id", mc.ShowRelation)
		relationsGroup.POST("/:id", mc.DeleteRelation)
		relationsGroup.GET("", mc.ShowAllRelations)
		relationsGroup.POST("", mc.InsertRelation)
	}

	// Group for "/fields" routes
	fieldsGroup := s.e.Group("/fields")
	{
		fieldsGroup.GET("/:id", mc.ShowField) // :id is the entityname where the field belongs
		fieldsGroup.POST("", mc.InsertField)
		fieldsGroup.POST("/:id", mc.DeleteField) // :id is the entityname where the field belongs to
	}

	// Group for "/lookups" routes
	lookupsGroup := s.e.Group("/lookups")
	{
		lookupsGroup.GET("", mc.ShowAllLookups)
		lookupsGroup.POST("", mc.InsertLookup)
		lookupsGroup.GET("/:id", mc.ShowLookup)
		lookupsGroup.POST("/:id", mc.ModifyLookup) // save, delete Lookup
	}

}

func (s GuiServer) StartServer(port int) {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprint(":", port)))
}
