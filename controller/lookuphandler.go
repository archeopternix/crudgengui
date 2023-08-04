package controller

import (
	"crudgengui/model"
	"net/http"

	"sync"
  "strconv"
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

  // delete lookup from list
  del := c.FormValue("delete")  
  if del =="true" { 
    fmt.Println("delete parameter: ",del)
    
  	if err := mc.repo.DeleteLookup(id); err != nil {
  		return echo.NewHTTPError(http.StatusNotFound, err.Error())
  	}
  } 

  // delete text from lookup 
  deltx := c.FormValue("deletetext")  
  if deltx =="true" { 
    fmt.Println("delete Txt parameter: ",deltx)
    
  	l, ok := mc.repo.GetLookup(id)
    if !ok {
  		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup to add entry '%v' not found", id))
  	}
    idx, err:= strconv.Atoi(c.FormValue("index"))
    if err!=nil {
      return echo.NewHTTPError(http.StatusNotFound, err)  
    }

    l.Delete(idx)
    if err:= mc.repo.SaveOrUpdateLookup(id,l); err!=nil {
      return echo.NewHTTPError(http.StatusNotFound, err.Error())  
    }
  } 

  // add attribute to lookup
  add := c.FormValue("add")
  if add =="true" { 
    fmt.Println("add parameter: ",add)
    
  	l, ok := mc.repo.GetLookup(id)
    if !ok {
  		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup to add entry '%v' not found", id))
  	}
    l.Add(c.FormValue("lookup_text"))
    if err:= mc.repo.SaveOrUpdateLookup(id,l); err!=nil {
      return echo.NewHTTPError(http.StatusNotFound, err.Error())  
    }
  }

  // in any case show the list of lookups
	m := mc.repo.GetModel()
	return c.Render(http.StatusOK, "lookups.html", map[string]interface{}{
		"model": m,
		"title": "Lookups",
	})
}

// ShowLookup /lookups/:id
func (mc ModelController) ShowLookup(c echo.Context) error {
   // Entity id from path `/lookups/:id`
	id := c.Param("id")
	lookup, ok := mc.repo.GetLookup(id)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup '%v' not found", id))
	}
	return c.Render(http.StatusOK, "lookup.html", map[string]interface{}{
		"model": lookup,
		"title": id,
	})
}