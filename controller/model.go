package controller

import (
	"crudgengui/model"
	"fmt"
	"os"
	"strings"
)

var m *model.Model
var yamlfile string

func SaveOrUpdateRelation(r *model.Relation) error {
	m.Relations[strings.ToLower(r.Name)] = *r

	err := saveYaml()
	return err
}

func DeleteRelation(name string) error {
	delete(m.Relations, strings.ToLower(name))

	err := saveYaml()
	return err
}

func GetAllRelations() map[string]model.Relation {
	return m.Relations
}

func GetRelation(name string) (model.Relation, bool) {
	r, ok := m.Relations[strings.ToLower(name)]
	return r, ok
}

func SaveOrUpdateEntity(e *model.Entity) error {
	m.Entities[strings.ToLower(e.Name)] = *e

	err := saveYaml()
	return err
}

func DeleteEntity(name string) error {
	delete(m.Entities, strings.ToLower(name))

	err := saveYaml()
	return err
}

func GetAllEntities() map[string]model.Entity {
	return m.Entities
}

func GetEntity(name string) (model.Entity, bool) {
	r, ok :=  m.Entities[strings.ToLower(name)]
  return r, ok 
}

// --------------  field ---------------
func SaveOrUpdateField(ename string, f *model.Field) error {
  e, ok:=GetEntity(strings.ToLower(ename))
  if !ok {
    return fmt.Errorf("Entity '%s' not found",ename)
  }
  e.Fields[strings.ToLower(f.Name)]=*f
  
	m.Entities[strings.ToLower(ename)] = e

	err := saveYaml()
	return err
}

func DeleteField(ename string,fname string) error {
  e, ok:=GetEntity(strings.ToLower(ename))
  if !ok {
    return fmt.Errorf("Entity '%s' not found",ename)
  }
  
	delete(e.Fields, strings.ToLower(fname))
  m.Entities[strings.ToLower(ename)] = e

	err := saveYaml()
	return err
}

func GetAllFields(ename string) (map[string]model.Field, error) {
  e, ok:=GetEntity(strings.ToLower(ename))
  if !ok {
    return nil, fmt.Errorf("Entity '%s' not found",ename)
  }
	return e.Fields, nil
}


func init() {
	m = model.NewModel()
}

func GetModel() *model.Model {
	return m
}

func saveYaml() error {
	file, err := os.Create(yamlfile)
	if err != nil {
		return err
	}
	err = m.WriteYAML(file)
	if err != nil {
		return err
	}
	return nil
}

func LoadModel(name string) (err error) {
	yamlfile = name
	file, err := os.Open(yamlfile)
	if err != nil {
		return err
	}
	err = m.ReadYAML(file)
	if err != nil {
		return err
	}
	return nil
}
