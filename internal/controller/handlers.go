// Package controller implements the functionalities used by the router (based on echo)
// and accesses the repository
package controller

import (
	repository "crudgengui/internal/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequestData is the standardized struct to hold data, additional text fields
// and error messages for submission to the golang template based webpages
type RequestData struct {
	Text  map[string]string // 'title': page title + 'menu': selected menu entry
	Data  map[string]interface{}
	Error map[string]string // form field  reference and error text
}

func newRequestData(text map[string]string, data map[string]interface{}) *RequestData {
	rd := new(RequestData)
	rd.Text = text
	rd.Data = data
	rd.Error = make(map[string]string)
	return rd
}

func (rd *RequestData) addError(key string, value string) error {
	if _, ok := rd.Error[key]; ok {
		return fmt.Errorf("key '%v' already exists in errors", key)
	}
	rd.Error[key] = value
	return nil
}

// ModelController attaches a concrete model implementation to the controller logic.
// repository/ModelRepositiory is the interface that has to satisfied to act as a valid
// repository
type ModelController struct {
	repo *repository.ModelRepository
}

// NewModelController creates a new ModelController instance and assignes a concrete
// model implementation to access a repository which persists the data
func NewModelController(r *repository.ModelRepository) *ModelController {
	mc := new(ModelController)
	mc.repo = r
	return mc
}

// ShowDashboard shows the startpage (index.html)
func (mc ModelController) ShowDashboard(c echo.Context) error {
	m := mc.repo.GetModel()
	text := map[string]string{
		"title":         "Home",
		"menu":          "menu_home",
		"entitycount":   fmt.Sprint(len(m.Entities)),
		"relationcount": fmt.Sprint(len(m.Relations)),
		"lookupcount":   fmt.Sprint(len(m.Lookups)),
	}
	rd := newRequestData(text, map[string]interface{}{
		"model": m,
	})
	return c.Render(http.StatusOK, "index.html", rd)
}
