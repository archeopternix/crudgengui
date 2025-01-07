// Package internal GenerationDSL project main.go
package repository

import (
	"fmt"
	"io/ioutil"

	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	model "crudgengui/internal/model"
	"crudgengui/pkg"

	"github.com/gertd/go-pluralize"
	"gopkg.in/yaml.v3"
)

// Generator for file generation, holds all Modules
type Generator struct {
	Modules map[string]Module
}

// Module is one independent dedicated functional unit that holds all Tasks
// (activities) to generate a certain part of an application (e.g. HTML view, Entities)
type Module struct {
	path  string
	Name  string `yaml:"name"`
	Tasks []Task
}

// Task is a single task for file generation which could be the copy of file or
// the generation based on template execution.
//
// Currently 2 modes are supported 'copy' or 'template'.
// Appdate = true indicates that the whole Application structure is submitted to
// the template generator. When Filename is set (not nil) the whole Application
// will be send to the template execution. If Filename is empty the generator
// iterates over all Entities and calls the template generator with a single entity.
// Filename is provided without file extension
type Task struct {
	Kind     string   `yaml:"kind"` // currently supported: copy, template
	Source   []string `yaml:"source"`
	Target   string   `yaml:"target"`             // target directory - filename wil be calculated
	Template string   `yaml:"template,omitempty"` // name of the template from template file
	Fileext  string   `yaml:"fileext,omitempty"`  // file extension for the generated file
	Filename string   `yaml:"filename,omitempty"` // when Filename is set (not nil) the whole Application will be send to the template execution
}

// Type returns the type of task based on the Kind and whether Filename is populated
func (t *Task) Type() string {
	if t.Kind == "copy" {
		return "copy"
	}
	if t.Kind == "template" {
		if t.Filename != "" {
			return "singletemplate"
		}
		return "multitemplate"
	}
	if (t.Kind == "singletemplate") || (t.Kind == "multitemplate") {
		return t.Kind
	}
	return "unknown"
}

var once sync.Once
var generator *Generator

// NewGenerator creates a singleton Generator or returns the pointer to an existing one
func NewGenerator() *Generator {
	// call the creation exactly one
	once.Do(func() {
		generator = new(Generator)
		generator.Modules = make(map[string]Module)
	})

	return generator
}

// AddModules reads in 1..n 'Module' from an YAML file and adds it to the generator configuration
// In a post processing step source/target filenames and filepaths will be cleaned
func (c *Generator) AddModules(filenames ...string) error {
	for _, filename := range filenames {
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("YAML file %v could not be loaded: #%v ", filename, err)
		}

		var m Module

		err = yaml.Unmarshal(yamlFile, &m)
		if err != nil {
			return fmt.Errorf("YAML file %v could not be unmarshalled: #%v ", filename, err)
		}

		m.path = filepath.Dir(filename)
		var tasks []Task
		for _, t := range m.Tasks {
			tasks = append(tasks, t.CleanPaths(m.path))
		}
		m.Tasks = tasks

		c.Modules[m.Name] = m
		slog.Debug(fmt.Sprintf("read in module '%s' from file '%s'", m.Name, filename))
	}
	return nil
}

// GenerateAll calls the GenerateModule function of all Modules
func (c Generator) GenerateAll(app *model.Model, genpath string) error {
	for _, m := range c.Modules {
		if err := m.GenerateModule(app, genpath); err != nil {
			return err
		}
	}
	return nil
}

// GenerateModule generates a 'Module' based on the Generator configuration.
// Currently implemented Modules are:
//
// kind: copy:
// - source: contains all the source files that will be copied 1:1
// - target: is the path where all source files will be copied into
//
// kind: template:
// - source: contains all the template files that will be used
// - target: is the path where all source files will be copied into
// - template: name of the primary template used for generation {{define "kinds"}}
// - fileext: is the extension of the generated files
// - filename (optional): when ilename is set (not nil) the whole Application will be send to the template execution
func (m *Module) GenerateModule(app *model.Model, genpath string) error {
	// pluralize and singularize functions for templates
	pl := pluralize.NewClient()
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		"lowercase": strings.ToLower, "title": strings.Title, "uppercase": strings.ToUpper, "singular": pl.Singular, "plural": pl.Plural, "inc": func(counter int) int { return counter + 1 },
	}

	for _, t := range m.Tasks {
		// check or create path
		path := filepath.Join(genpath, app.Name, t.Target)
		if err := pkg.CheckMkdir(path); err != nil {
			if _, ok := err.(*pkg.DirectoryExistError); !ok {
				return err
			}
		}

		switch t.Type() {
		case "copy":
			// copying all files from .Source to .Target
			for _, src := range t.Source {
				path := filepath.Join(genpath, app.Name, t.Target, filepath.Base(src))
				if err := pkg.CopyFile(src, path); err != nil {
					if _, ok := err.(*pkg.FileExistError); !ok {
						return err
					}
				}
				slog.Debug("file created", "file", path)
			}

		case "singletemplate":
			// Create a template, add the function map, and parse the text.
			tmpl, err := template.New(t.Template).Funcs(funcMap).ParseFiles(t.Source...)
			if err != nil {
				return fmt.Errorf("creating template: %v", err)
			}

			file := filepath.Join(genpath, app.Name, t.Target, strings.ToLower(t.Filename)+t.Fileext)
			writer, err := os.Create(file)
			if err != nil {
				return fmt.Errorf("creating file %v", err)
			}
			defer writer.Close()
			if err := tmpl.ExecuteTemplate(writer, t.Template, app); err != nil {
				return fmt.Errorf("executing template %v", err)
			}
			slog.Debug("template created", "template", t.Template, "file", file)

		case "multitemplate":
			// Create a template, add the function map, and parse the text.
			tmpl, err := template.New(t.Template).Funcs(funcMap).ParseFiles(t.Source...)
			if err != nil {
				return fmt.Errorf("creating template: %v", err)
			}

			for _, entity := range app.Entities {
				file := filepath.Join(genpath, app.Name, t.Target, strings.ToLower(entity.Name)) + t.Fileext
				writer, err := os.Create(file)
				if err != nil {
					return fmt.Errorf("creating file %v", err)
				}
				defer writer.Close()
				type DataStruct struct {
					Entity    model.Entity
					Lookups   map[string]model.Lookup
					AppName   string
					TimeStamp string
				}
				entityStruct := DataStruct{
					Entity:    entity,
					AppName:   app.Name,
					TimeStamp: app.TimeStamp(),
				}
				entityStruct.Lookups = make(map[string]model.Lookup)
				entityStruct.Lookups = app.Lookups

				if err := tmpl.ExecuteTemplate(writer, t.Template, entityStruct); err != nil {
					return fmt.Errorf("executing template %v", err)
				}
				slog.Debug("Template created", "template", t.Template, "entity", entity.Name, "file", file)
			}

		default:
			return fmt.Errorf("unknown generator operation '%s'", t.Kind)
		}
	}

	slog.Debug("generated successful", "module", m.Name)
	return nil
}

// AddTask adds a task to a module an cleans the target and source path
func (m *Module) AddTask(t *Task) {
	m.Tasks = append(m.Tasks, t.CleanPaths(m.path))
}

// CleanPaths cleans the target and source path and adds to the sourcepath the
// filepath of the module
// - source: the module path will be added so fields are accessible from root of application
// - target: just clean the target path
func (t *Task) CleanPaths(modulepath string) Task {

	task := *t
	task.Source = nil
	for _, s := range t.Source {
		task.Source = append(task.Source, filepath.Join(modulepath, filepath.Clean(s)))
	}
	task.Target = filepath.Clean(t.Target)
	return task
}

/*
// LoadFromFile reads in the full generator configuration from an YAML file
// In a post processing step for all loaded module the source/target filenames
// and filepaths are cleaned
func (c *Generator) LoadFromFile(filename string) error {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be loaded: #%v ", filename, err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be unmarshalled: #%v ", filename, err)
	}

	for key, m := range c.Modules {
		m.path = filepath.Dir(filename)
		var tasks []Task
		for _, t := range m.Tasks {
			tasks = append(tasks, t.CleanPaths(m.path))
		}
		m.Tasks = tasks
		c.Modules[key] = m
	}

	log.Printf("read in Generator configuration from file '%s'", filename)
	return nil
}

// SaveToFile saves the full generator configuration to a YAML file
func (c Generator) SaveToFile(filename string) error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be marshalled: #%v ", filename, err)
	}

	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		return fmt.Errorf("YAML file %v could not be saved: #%v ", filename, err)
	}

	log.Printf("saved Generator configuration to file '%s'", filename)
	return nil
}

*/
