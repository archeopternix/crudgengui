package internal

import (
	"html/template"
	"log"
	"strings"

	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

// TemplateRegistry defines the template registry struct
// Usage with echo:
//      templates := controller.NewTemplateRegistry()
//      e.Renderer = templates
type TemplateRegistry struct {
	templates map[string]*template.Template
	funcMap   template.FuncMap
}

// Render implements e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return fmt.Errorf("Template not found -> %s", name)
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

// getFuncMap defines useful functions for templates
func (t *TemplateRegistry) getFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"title":     strings.Title,
		"lowercase": strings.ToLower,
		"uppercase": strings.ToUpper,
	}
	return funcMap
}

// NewTemplateRegistry creates a new template registry
func NewTemplateRegistry() *TemplateRegistry {
	tr := new(TemplateRegistry)
	tr.funcMap = tr.getFuncMap()
	tr.templates = make(map[string]*template.Template)

	return tr
}

// AddTemplate adds another template into the registry which could referenced by its name and could have a base template. To define a base template keep 2nd parameter empty ""
func (tr *TemplateRegistry) addTemplate(name string, basetemplate string, filenames ...string) error {
	var err error
	if len(name) < 1 {
		return fmt.Errorf("Template name must be defined")
	}
	if len(basetemplate) < 1 {
		// base template
		log.Println("New base template:", name, ", ", basetemplate)
		tr.templates[name], err = template.New(name).Funcs(tr.funcMap).ParseFiles(filenames...)
		if err != nil {
			return fmt.Errorf("Parse base template '%s': %v", name, err)
		}
	} else {
		if _, ok := tr.templates[basetemplate]; !ok {
			return fmt.Errorf("No base template '%s' by adding individual template '%s'", basetemplate, name)
		}
		// individual templates
		if len(filenames) > 0 {
			log.Println("New template:", name, ", ", basetemplate)
			tr.templates[name], err = template.Must(tr.templates[basetemplate].Clone()).ParseFiles(filenames...)
			if err != nil {
				return fmt.Errorf("Parse individual template '%s' based on base template '%s': %v", name, basetemplate, err)
			}
		} else {
			log.Println("New template without files:", name, ", ", basetemplate)
			tr.templates[name], err = tr.templates[basetemplate].Clone()
			if err != nil {
				return fmt.Errorf("Create individual template '%s' based on base template '%s': %v", name, basetemplate, err)
			}
		}
	}
	return nil
}

func (tr *TemplateRegistry) AddTemplateOrPanic(name, base string, files ...string) {
	if err := tr.addTemplate(name, base, files...); err != nil {
		log.Panic(err)
	}
}
