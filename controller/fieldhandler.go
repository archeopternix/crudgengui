package controller

import (
	"net/http"
	"sync"
	model "crudgengui/model"
	"fmt"

	"github.com/labstack/echo/v4"
)

// ------------- Fields -------------

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
	return mc.ShowAllEntities(c)

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

  if (field.Length =="0") {
     field.Length ="" 
  }
    if (field.MaxLength =="0") {
     field.MaxLength ="" 
  }
  if (field.Size =="0") {
     field.Size ="" 
  }
  if (field.Max =="0") {
     field.Max ="" 
  }
  if (field.Step =="0") {
     field.Step ="" 
  }
  
	if (len(field.Name) < 3) || (len(ename) < 3) {
		return nil
	}

	if err := mc.repo.SaveOrUpdateField(ename, field); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

  // Show entity with the name: ename 
  return mc.showEntity(c,ename)
}

// ShowField shows detail page for Field
// /fields/:id shows detail page to create the Field definition. :id is the entity
// or
// /fields/:id?field=myfield shows detail page for the Field . :id is the entity
func (mc ModelController) ShowField(c echo.Context) error {
	
  // Entity id from path `/entities/:id`
	id := c.Param("id")
	entity, ok := mc.repo.GetEntity(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity '%v' not found", id))
	}
  
  // Query parameter ?field=myfield
  fieldname := c.QueryParam("field")
  if len(fieldname)>0 {   
    // Show field definition of entity
    field, ok := mc.repo.GetField(id,fieldname)
  	if !ok {
  		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Field '%v' in Entity '%v' not found", fieldname, id))
  	}  
    data := struct {
      Field   model.Field
      LookupNames []string
    }{
      *field,
       mc.repo.GetAllLookupNames(),
    }
  	return c.Render(http.StatusOK, "field.html", map[string]interface{}{
  		"model": data,
  		"entityname": entity.Name,
  		"title":  fmt.Sprint("Field: '",field.Name, "'"),
  	}) 
  } 
    data := struct {
      Field   model.Field
      LookupNames []string
    }{
       model.Field{},
       mc.repo.GetAllLookupNames(),
    }
  // Create a new field
    	return c.Render(http.StatusOK, "field.html", map[string]interface{}{
  		"model":  data,
  		"entityname": entity.Name,
  		"title":  fmt.Sprint("Create a new Field"),
  	})   
  
}