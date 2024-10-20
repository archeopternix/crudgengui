// Package model defines the entities and relations used as a bases for the CRUD generator
package model

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)


// Model holds all entities, relations and lookups and is able to persist them as a YAML file
type Model struct {
	Entities  map[string]Entity
	Relations map[string]Relation
  Lookups map[string]Lookup
}

// NewModel creates a new model and initialize the maps used for entities, relations and lookups
func NewModel() *Model {
	m := new(Model)
	m.Entities = make(map[string]Entity)
	m.Relations = make(map[string]Relation)
  m.Lookups = make(map[string]Lookup)
	return m
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
  for _, r:= range m.Relations {
    if r.ContainsEntity(entityname) {
      return true
    }
  }
  return false
}
