package controller

import (
	"net/http"
	"sync"
	model "crudgengui/model"
	"fmt"

	"github.com/labstack/echo/v4"
)

// ------------- Relations -------------
// ShowAllRelations
func (mc ModelController) ShowAllRelations(c echo.Context) error {
	m := mc.repo.GetModel()

  text:=map[string]string{
  "title": "Relations",
  "menu": "menu_relations",
	}
  rd:= newRequestData(text,map[string]interface{}{
		"model": m,
	})
  
  return c.Render(http.StatusOK, "relations.html", rd)  

}

// ShowRelation shows detail page to model or edit screen for new model
func (mc ModelController) ShowRelation(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// DeleteRelation shows detail page to model or edit screen for new model
func (mc ModelController) DeleteRelation(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if _, ok := mc.repo.GetRelation(id); !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	if err := mc.repo.DeleteRelation(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/relations")

}

// InsertRelation
func (mc ModelController) InsertRelation(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	r := new(model.Relation)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if (r.Source == "Please Select") || (r.Destination == "Please Select") || (r.Destination == r.Source) {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Source '%s' and/or destination '%s' Entity are conflicting", r.Source, r.Destination))
	}

	rname := r.Source + r.Type + r.Destination

	if len(rname) < 3 {
		return nil
	}

	if _, ok := mc.repo.GetRelation(rname); ok {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Name %v for Relation is already in use", rname))
	}

	if err := mc.repo.SaveOrUpdateRelation(rname, r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return mc.ShowAllEntities(c)
}
