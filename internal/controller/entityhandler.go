package controller

import (
	model "crudgengui/internal/model"
	"crudgengui/pkg"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

// ------------- Entities -------------
// showEntity helper function to show detail page for entity
func (mc ModelController) showEntity(c echo.Context, entityname string) error {
	entity, ok := mc.repo.GetEntity(entityname)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity '%v' not found", entityname))
	}

	text := map[string]string{
		"title": "Entity: " + strings.Title(entityname),
		"menu":  "menu_entities",
	}
	rd := newRequestData(text, map[string]interface{}{
		"entity": entity,
	})

	return c.Render(http.StatusOK, "entity.html", rd)
}

// ShowAllEntities retrieves all entities from repo and shows list page
// route: GET /entities
func (mc ModelController) ShowAllEntities(c echo.Context) error {
	m := mc.repo.GetModel()
	text := map[string]string{
		"title": "Entities",
		"menu":  "menu_entities",
	}
	rd := newRequestData(text, map[string]interface{}{
		"model": m,
	})

	return c.Render(http.StatusOK, "entities.html", rd)
}

// InsertEntity inserts one entity into the repo and returns to list page
// route: FORM POST /entities
func (mc ModelController) InsertEntity(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	entity := model.NewEntity()
	if err := c.Bind(entity); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate entity
	if err := mc.validateEntity(entity); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := mc.repo.SaveOrUpdateEntity(entity); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/entities")
}

// ShowEntity shows detail page for an Entity or if a query parameter is set the respective Field
// Option 1: /entities/:id shows detail page of Entity
// Option 2: /entities/:id?field=myfield shows detail page to edit the Field definition
func (mc ModelController) ShowEntity(c echo.Context) error {
	// Entity id from path `/entities/:id`
	id := c.Param("id")
	if _, ok := mc.repo.GetEntity(id); !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity '%v' not found", id))
	}

	// Query parameter ?field=myfield
	fieldname := c.QueryParam("field")
	if len(fieldname) > 0 {
		// show detail for field
		if err := mc.showField(c, id, fieldname); err != nil {
			return err
		}
		return nil
	}

	// Show entity with the name: id
	return mc.showEntity(c, id)
}

// DeleteEntity shows detail page to model or edit screen for new model
// route: POST /lookups/:id
func (mc ModelController) DeleteEntity(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	if _, ok := mc.repo.GetEntity(id); !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("ID %v not found", id))
	}

	if err := mc.repo.DeleteEntity(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/entities")
}

// validateEntity validates the entity's name and type
func (mc ModelController) validateEntity(e *model.Entity) error {
	if !pkg.IsFirstLetterUppercase(e.Name) {
		e.Name = strings.Title(e.Name)
	}
	if len(e.Name) < 3 || len(e.Type) < 3 {
		return fmt.Errorf("Name %v needs to be 3 characters minimum", e.Name)
	}
	if _, ok := mc.repo.GetEntity(e.Name); ok {
		return fmt.Errorf("Name %v is already in use", e.Name)
	}

	// TODO: Check for existing name in lookups
	/* if _, ok := mc.repo.GetEntity(e.Name); ok {
	       return fmt.Errorf("Name %v is already in use", e.Name)
	   }
	*/
	return nil
}
