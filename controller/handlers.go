package controller
import (

	"net/http"

	"sync"

	model "crudgengui/model"
  repository "crudgengui/repository"
	"fmt"


	"github.com/labstack/echo/v4"
)

type ModelController struct {
  repo *repository.ModelRepository
}

func NewModelController(r *repository.ModelRepository) *ModelController {
  mc := new(ModelController)
  mc.repo=r
  return mc
}


func (mc ModelController) ShowDashboard(c echo.Context) error {
	m := mc.repo.GetModel()
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"model": m,
		"title": "Home",
	})
}

// ------------- Entities -------------

// showAllEntities
func (mc ModelController) ShowAllEntities(c echo.Context) error {
	m := mc.repo.GetModel()
	return c.Render(http.StatusOK, "entities.html", map[string]interface{}{
		"model": m,
		"title": "Entities",
	})
}

// insertEntity
func (mc ModelController) InsertEntity(c echo.Context) error {
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
	if _, ok := mc.repo.GetEntity(e.Name); ok {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Name %v is already in use", e.Name))
	}

	if err := mc.repo.SaveOrUpdateEntity(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return mc.ShowAllEntities(c)
}

// showEntity shows detail page to model or edit screen for new model
func (mc ModelController) ShowEntity(c echo.Context) error {

	id := c.Param("id")
	entity, ok := mc.repo.GetEntity(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	all := mc.repo.GetAllEntities()
	return c.Render(http.StatusOK, "entity.html", map[string]interface{}{
		"model":  all,
		"entity": entity,
		"title":  fmt.Sprint("Entity: ", entity.Name),
	})
}

// deleteEntity shows detail page to model or edit screen for new model
func (mc ModelController) DeleteEntity(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if _, ok := mc.repo.GetEntity(id); !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	if err := mc.repo.DeleteEntity(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/entities")
}

// ------------- Relations -------------

// showAllRelations
func (mc ModelController) ShowAllRelations(c echo.Context) error {
	m := mc.repo.GetModel()
	return c.Render(http.StatusOK, "relations.html", map[string]interface{}{
		"model": m,
		"title": "Relations",
	})
}

// showRelation shows detail page to model or edit screen for new model
func (mc ModelController) ShowRelation(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// deleteRelation shows detail page to model or edit screen for new model
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

// insertRelation
func (mc ModelController) InsertRelation(c echo.Context) error {
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

	if _, ok := mc.repo.GetEntity(r.Name); ok {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("Name %v is already in use", r.Name))
	}

	if err := mc.repo.SaveOrUpdateRelation(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return mc.ShowAllEntities(c)
}

// ------------- Fields -------------

// deleteRelation shows detail page to model or edit screen for new model
func (mc ModelController) DeleteField(c echo.Context) error {
	// User ID from path `users/:id`
	fname := c.FormValue("field_name")
	ename := c.FormValue("entity_name")

	if (len(fname) < 2) || (len(ename) < 2) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity '%s', field '%s' not found", ename, fname))
	}

	if err := mc.repo.DeleteField(ename, fname); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return mc.ShowAllEntities(c)

}

// insertRelation
func (mc ModelController) InsertField(c echo.Context) error {
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

	if err := mc.repo.SaveOrUpdateField(ename,field); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return mc.ShowAllEntities(c)
}
