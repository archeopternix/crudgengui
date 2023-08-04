package model

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type Relation struct {
	Name        string `json:"name" form:"relation-name"`
	Type        string `json:"type" form:"relation-type"` // 'One-to-Many' | 'Many-to-Many'
	Source      string `json:"source" form:"relation-source"`
	Destination string `json:"destination" form:"relation-destination"`
}

func (r Relation) ContainsEntity(name string) bool {
  if (r.Source == name) || (r.Destination == name) {
    return true
  }
  return false
}

// Lookup is a string list
type Lookup struct {
  List []string
}

func NewLookup() *Lookup {
  return new(Lookup)
}

func (f *Lookup)Add(text string){
  f.List=append(f.List,text)
}

func (f *Lookup)Delete(i int){
  copy(f.List[i:], f.List[i+1:]) // Shift a[i+1:] left one index.
  f.List[len(f.List)-1] = ""     // Erase last element (write zero value).
  f.List = f.List[:len(f.List)-1]     // Truncate slice.
}

type Model struct {
	Entities  map[string]Entity
	Relations map[string]Relation
  Lookups map[string]Lookup
}

func NewModel() *Model {
	m := new(Model)
	m.Entities = make(map[string]Entity)
	m.Relations = make(map[string]Relation)
  m.Lookups = make(map[string]Lookup)
	return m
}

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
