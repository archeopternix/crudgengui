// Package model defines the entities and relations used as a bases for the CRUD generator
package model

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gertd/go-pluralize"
	"gopkg.in/yaml.v3"
)

// Model holds all entities, relations and lookups and is able to persist them as a YAML file
type Model struct {
	Name string `yaml:"name"`
	Settings
	Entities  map[string]Entity
	Relations map[string]Relation
	Lookups   map[string]Lookup
}

// Settings is the definition of the global attributes
type Settings struct {
	CurrencySymbol    string `yaml:"currency_symbol"`
	DecimalSeparator  string `yaml:"decimal_separator"`
	ThousendSeparator string `yaml:"thousend_separator"`
	TimeFormat        string `yaml:"time_format"`
	DateFormat        string `yaml:"date_format"`
}

func NewEUROSettings() Settings {
	s := Settings{
		CurrencySymbol:    `€`,
		DecimalSeparator:  `,`,
		ThousendSeparator: `.`,
		TimeFormat:        "15:04:05",
		DateFormat:        "02.01.2006",
	}
	return s
}

func NewUSSettings() Settings {
	s := Settings{
		CurrencySymbol:    `$`,
		DecimalSeparator:  `.`,
		ThousendSeparator: `,`,
		TimeFormat:        "15:04:05",
		DateFormat:        "01/02/2006",
	}
	return s
}

// NewModel creates a new model and initialize the maps used for entities, relations and lookups
func NewModel() *Model {
	m := new(Model)
	m.Settings = NewEUROSettings()
	m.Entities = make(map[string]Entity)
	m.Relations = make(map[string]Relation)
	m.Lookups = make(map[string]Lookup)
	return m
}

// TimeStamp needed for file generation. Will be added in the header of each file
// to track the creation date and time of each file
func (m *Model) TimeStamp() string {
	return time.Now().Format(m.DateFormat + " " + m.TimeFormat)
}

// ReadYAML reads the model as YAML from an io.Reader
func (m *Model) ReadYAML(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error in model read, %v", err)
	}
	err = yaml.Unmarshal([]byte(data), m)
	if err != nil {
		return fmt.Errorf("error in model read, %v", err)
	}

	if m.Relations == nil {
		m.Relations = make(map[string]Relation)
	}
	if m.Entities == nil {
		m.Entities = make(map[string]Entity)
	}
	return nil
}

// WriteYAML writes the model as YAML fiel to an io.Writer
func (m *Model) WriteYAML(writer io.Writer) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("error in model write, %v", err)
	}
	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("error in model write, %v", err)
	}
	return nil
}

// EntityInRealtions checks if the entityname is contained in Source
// or Destination of one of the entities
func (m Model) EntityInRealtions(entityname string) bool {
	for _, r := range m.Relations {
		if r.ContainsEntity(entityname) {
			return true
		}
	}
	return false
}

// creates a deep copy of the Model
func (om *Model) deepCopy() *Model {
	copy := &Model{
		Name:      om.Name,
		Settings:  om.Settings,
		Entities:  make(map[string]Entity),
		Relations: make(map[string]Relation),
		Lookups:   make(map[string]Lookup),
	}
	for key, entity := range om.Entities {
		copy.Entities[key] = entity
	}
	for key, relations := range om.Relations {
		copy.Relations[key] = relations
	}
	for key, lookups := range om.Lookups {
		copy.Lookups[key] = lookups
	}
	return copy
}

// parseDependencies parse all entities for lookup fields, add unique ID field
// and parse relations between entities and therefore adds dedicated fields for
// parent/child relations and scans for lookups and parent-child relationships
// and therefore creates necessary additional entities (e.g. lookup entities)
// or add additional fields (e.g. Id field for every entity)
func (m Model) ParseDependencies() (*Model, error) {
	copy := m.deepCopy()

	for key, entity := range copy.Entities {
		// Parse fields
		for i := range entity.Fields {
			field := &entity.Fields[i]

			// If a lookup field is present check for lookup table
			if field.Type == "Lookup" {
				lookup := strings.ToLower(field.Lookup)
				if _, ok := copy.Lookups[lookup]; !ok {
					return nil, fmt.Errorf("lookup with name '%s' could not be found", field.Lookup)
				}
				field.Object = lookup
			}
		}

		// Add an ID field to all entities if not yet exists
		if idx := entity.GetFieldIndexByName("ID"); idx < 0 {
			entity := copy.Entities[key]
			entity.Add(Field{Name: "ID", Type: "Integer", Required: true, Auto: true})
			copy.Entities[key] = entity
		}
	}

	// Add fields for relationships between entities
	// Source .. Parent
	// Destination .. Child
	for _, relation := range copy.Relations {
		sourcename := strings.ToLower(relation.Source)
		destname := strings.ToLower(relation.Destination)

		if relation.Type == "One-to-Many" {
			// Add Child to Parent field
			dest := copy.Entities[destname]
			fieldname := strings.Title(sourcename) + "_ID"
			if idx := dest.GetFieldIndexByName(fieldname); idx < 0 {
				dest.Add(Field{Name: fieldname, Type: "Child", Object: sourcename, Auto: true}) // it is the child field
				copy.Entities[destname] = dest
			}

			// Add Childs to Parent field
			plural := pluralize.NewClient()
			source := copy.Entities[sourcename]
			fieldname = plural.Plural(strings.Title(destname))
			if idx := dest.GetFieldIndexByName(fieldname); idx < 0 {
				source.Add(Field{Name: fieldname, Type: "Parent", Object: destname, Auto: true}) // it is the parent field
				copy.Entities[sourcename] = source
			}
		}
	}
	/*
		for _,lookup := range a.Lookups {
			// do nothing so far
		}
	*/

	return copy, nil
}

func WriteToFile(m *Model, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = m.WriteYAML(file)
	if err != nil {
		return err
	}
	return nil
}
