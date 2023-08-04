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
	mc.repo = r
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

// ShowEntity shows detail page for an Entity or if a query parameter is set the respective Field
// Option 1: /entities/:id shows detail page of Entity
// Option 2: /entities/:id?field=myfield shows detail page to edit the Field definition
func (mc ModelController) ShowEntity(c echo.Context) error {
	
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
  	return c.Render(http.StatusOK, "field.html", map[string]interface{}{
  		"model":  field,
  		"entityname": entity.Name,
  		"title":  fmt.Sprint("Field: '",field.Name, "'"),
  	}) 
  } else {
    
    // Show entity
  	all := mc.repo.GetAllEntities()
  	return c.Render(http.StatusOK, "entity.html", map[string]interface{}{
  		"model":  all,
  		"entity": entity,
  		"title":  fmt.Sprint("Entity: ", entity.Name),
  	})  
  }
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

// ShowAllRelations
func (mc ModelController) ShowAllRelations(c echo.Context) error {
	m := mc.repo.GetModel()
	return c.Render(http.StatusOK, "relations.html", map[string]interface{}{
		"model": m,
		"title": "Relations",
	})
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

  entity, ok := mc.repo.GetEntity(ename)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Entity %v not found", ename))
	}

	all := mc.repo.GetAllEntities()
	return c.Render(http.StatusOK, "entity.html", map[string]interface{}{
		"model":  all,
		"entity": entity,
		"title":  fmt.Sprint("Entity: ", entity.Name),
	})
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
  	return c.Render(http.StatusOK, "field.html", map[string]interface{}{
  		"model":  field,
  		"entityname": entity.Name,
  		"title":  fmt.Sprint("Field: '",field.Name, "'"),
  	}) 
  } 

  // Create a new field
    	return c.Render(http.StatusOK, "field.html", map[string]interface{}{
  		"model":  model.Field{},
  		"entityname": entity.Name,
  		"title":  fmt.Sprint("Create a new Field"),
  	})   
  
}