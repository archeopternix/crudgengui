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
// showLookup helper function to show lookup based on the name of the lookup
func (mc ModelController) showLookup(c echo.Context, lookupname string) error {
	lookup, ok := mc.repo.GetLookup(lookupname)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup '%v' not found", lookupname))
	}

  text:=map[string]string{
		"title": lookupname,
    "menu": "menu_lookups",
	}
  rd:= newRequestData(text,map[string]interface{}{
		"lookup": lookup,
	})
	return c.Render(http.StatusOK, "lookup.html", rd)
}

// ShowAllLookups
func (mc ModelController) ShowAllLookups(c echo.Context) error {
	m := mc.repo.GetModel()
   text:=map[string]string{
		"title": "Lookups",
    "menu": "menu_lookups",
	}
  rd:= newRequestData(text,map[string]interface{}{
		"lookups": m.Lookups,
	})
	return c.Render(http.StatusOK, "lookups.html", rd)
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
  	if err := mc.repo.DeleteLookup(id); err != nil {
  		return echo.NewHTTPError(http.StatusNotFound, err.Error())
  	}
  } 

  // delete text from lookup 
  deltx := c.FormValue("deletetext")  
  if deltx =="true" { 
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
    return mc.showLookup(c,id)
  } 

  // add attribute to lookup
  add := c.FormValue("add")
  if add =="true" { 
  	l, ok := mc.repo.GetLookup(id)
    if !ok {
  		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("Lookup to add entry '%v' not found", id))
  	}
    l.Add(c.FormValue("lookup_text"))
    if err:= mc.repo.SaveOrUpdateLookup(id,l); err!=nil {
      return echo.NewHTTPError(http.StatusNotFound, err.Error())  
    }
     return mc.showLookup(c,id)
  }

  return c.Redirect(http.StatusSeeOther,"/lookups") 
}

// ShowLookup /lookups/:id
func (mc ModelController) ShowLookup(c echo.Context) error {
   // Entity id from path `/lookups/:id`
	id := c.Param("id")
	
	return mc.showLookup(c, id)
}

