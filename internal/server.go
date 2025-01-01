package internal

import (
	controller "crudgengui/internal/controller"
	repository "crudgengui/internal/repository"
	"fmt"
	"log"
	"os"

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
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	// Create templates
	templates := NewTemplateRegistry()

	// base template
	if err := templates.AddTemplate("base.html", "", tpl+"base/side_navigation.html", tpl+"base/delete_popup.html", tpl+"base/base.html", tpl+"base/side_navigation.html", tpl+"base/top_navigation.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("entities.html", "base.html", tpl+"entity_popup.html", tpl+"base/script.html", tpl+"entities.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("lookups.html", "base.html", tpl+"lookup_popup.html", tpl+"base/script.html", tpl+"lookups.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("lookup.html", "base.html", tpl+"lookupadd_popup.html", tpl+"base/script.html", tpl+"lookup.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("relations.html", "base.html", tpl+"relation_popup.html", tpl+"base/script.html", tpl+"relations.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("entity.html", "base.html", tpl+"base/script.html", tpl+"entity.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("field.html", "base.html", tpl+"field.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("index.html", "base.html", tpl+"index.html"); err != nil {
		log.Panic(err)
	}
	if err := templates.AddTemplate("project.html", "base.html", tpl+"project.html"); err != nil {
		log.Panic(err)
	}

	s.e.Renderer = templates
}

func (s *GuiServer) setRoutes() {
	// Create Repository
	mc := controller.NewModelController(repository.NewModelRepository(repository.NewYAMLModel(base + "data/model.yaml")))

	s.e.GET("/", mc.ShowDashboard)
	s.e.GET("/project", mc.ShowProject)
	s.e.POST("/project", mc.SaveProject)

	s.e.GET("/entities/:id", mc.ShowEntity)
	s.e.POST("/entities/:id", mc.DeleteEntity)
	s.e.GET("/entities", mc.ShowAllEntities)
	s.e.POST("/entities", mc.InsertEntity)

	s.e.GET("/relations/:id", mc.ShowRelation)
	s.e.POST("/relations/:id", mc.DeleteRelation)
	s.e.GET("/relations", mc.ShowAllRelations)
	s.e.POST("/relations", mc.InsertRelation)

	s.e.GET("/fields/:id", mc.ShowField) // :id is the entityname where the field belongs
	s.e.POST("/fields", mc.InsertField)
	s.e.POST("/fields/:id", mc.DeleteField) // :id is the entityname where the field belongs to

	s.e.GET("/lookups", mc.ShowAllLookups)
	s.e.POST("/lookups", mc.InsertLookup)
	s.e.GET("/lookups/:id", mc.ShowLookup)
	s.e.POST("/lookups/:id", mc.ModifyLookup) // save, delete Lookup

}

func (s GuiServer) StartServer(port int) {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprint(":", port)))
}
