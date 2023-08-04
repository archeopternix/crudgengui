package controller

import (
	"crudgengui/model"
	"net/http"

	"sync"

	"fmt"

	"github.com/labstack/echo/v4"
)

// ------------- Lookups -------------

// ShowAllLookups
func (mc ModelController) ShowAllLookups(c echo.Context) error {
	m := mc.repo.GetModel()
	return c.Render(http.StatusOK, "lookups.html", map[string]interface{}{
		"model": m,
		"title": "Lookups",
	})
}


// InsertLookup
func (mc ModelController) InsertLookup(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	name := c.FormValue("lookup_name")
	if (len(name) < 2)  {
		return nil
	}
	if _, ok := mc.repo.GetLookup(name); ok {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Name %v is already in use", name))
	}

	if err := mc.repo.SaveOrUpdateLookup(name, model.NewLookup()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return mc.ShowAllLookups(c)
}

// ModifyLookup shows detail page to model or edit screen for new model
func (mc ModelController) ModifyLookup(c echo.Context) error {
	// Lookup id from path `lookups/:id`
	id := c.Param("id")
  if _, ok := mc.repo.GetLookup(id); !ok {
  	return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup '%v' not found", id))
  }
  
  del := c.FormValue("delete")
  fmt.Println("delete parameter: ",del)
  if del =="true" { 
    // delete attribute is set
  	if err := mc.repo.DeleteLookup(id); err != nil {
  		return echo.NewHTTPError(http.StatusNotFound, err.Error())
  	}
  
  	return c.Redirect(http.StatusTemporaryRedirect, "/lookups")
  } 
  
  return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup 'update' not implemented for '%v'", id))
}