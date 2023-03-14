package controller

import (
  "crudgengui/model"
  "os"
)

var m *model.Model
var yamlfile string


func SaveOrUpdateRelation(r *model.Relation) error {
  m.Relations[r.Name] = *r
  
  err:= saveYaml()
  return err
}

func DeleteRelation(name string) error{
  delete(m.Relations, name)

  err:= saveYaml()
  return err
}

func GetAllRelations() map[string]model.Relation{
  return m.Relations
}

func SaveOrUpdateEntity(e *model.Entity) error { 
  m.Entities[e.Name] = *e

  err:= saveYaml()
  return err
}

func DeleteEntity(name string) error {
  delete(m.Entities, name)

  err:= saveYaml()
  return err
}

func GetAllEntities() map[string]model.Entity{
  return m.Entities
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
		return  err
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
		return  err
	}
  return nil
}