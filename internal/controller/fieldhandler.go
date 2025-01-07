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

// ------------- Fields -------------
// showField helper function to show detail page for field
func (mc ModelController) showField(c echo.Context, entityname string, fieldname string) error {
	var rd *RequestData
	field := &model.Field{}
	titletext := "Create a new field: "

	// if filename defined show field detail page
	if len(fieldname) > 0 {
		var ok bool
		// show field detail page
		field, ok = mc.repo.GetField(entityname, fieldname)
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Field '%v' in Entity '%v' not found", fieldname, entityname))
		}
		titletext = "Field: " + fieldname
	}

	text := map[string]string{
		"title":      titletext,
		"menu":       "menu_entities",
		"entityname": strings.Title(entityname),
	}

	rd = newRequestData(text, map[string]interface{}{
		"field":       field,
		"lookupnames": mc.repo.GetAllLookupNames(),
	})

	return c.Render(http.StatusOK, "field.html", rd)
}

// DeleteField shows detail page to model or edit screen for new model
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

	return mc.showEntity(c, ename)
}

// InsertField
func (mc ModelController) InsertField(c echo.Context) error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	ename := c.FormValue("entity_name")

	field := new(model.Field)
	if err := c.Bind(field); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	field.Id = pkg.CleanID(field.Name)
	if field.Decimals == "0" {
		field.Decimals = ""
	}
	if field.Min == "0" {
		field.Min = ""
	}
	if field.Size == "0" {
		field.Size = ""
	}
	if field.Max == "0" {
		field.Max = ""
	}
	if field.Step == "0" {
		field.Step = ""
	}

	if (len(field.Name) < 3) || (len(ename) < 3) {
		return nil
	}

	if err := mc.repo.SaveOrUpdateField(ename, field); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Show entity with the name: ename
	return mc.showEntity(c, ename)
}

// ShowField shows detail page for Field
// /fields/:id shows detail page to create the Field definition. :id is the entity
// or
// /fields/:id?field=myfield shows detail page for the Field . :id is the entity
func (mc ModelController) ShowField(c echo.Context) error {

	// Entity id from path `/entities/:id`
	id := c.Param("id")
	_, ok := mc.repo.GetEntity(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity '%v' not found", id))
	}

	// Query parameter ?field=myfield
	fieldname := c.QueryParam("field")

	return mc.showField(c, id, fieldname)
}
