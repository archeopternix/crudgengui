{{define "mockrepo" -}}
// Package mockdatabase contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package database

import (
	"fmt"
	model "{{.AppName}}/model"
)

{{with .Entity}}

// {{.CleanName}}Repo is the interface for a {{.CleanName}} repository that will persist 
// and retrieve data and has to be implemented for concrete Databases 
// (e.g. db *sqlx.DB) or other respositories
type {{.CleanName}}Repo struct{
	data map[uint64]model.{{.CleanName}}
	count uint64
}

var {{.CleanName | lowercase}}repo *{{.CleanName}}Repo


// New{{.Name}}Repo creates a new instance of {{.CleanName}}Repo and initializes the 
// map and sets counter for ID to 1
func New{{.Name}}Repo() *{{.CleanName}}Repo {
	{{.CleanName | lowercase}}repo = new({{.CleanName}}Repo)
	{{.CleanName | lowercase}}repo.data = make(map[uint64]model.{{.CleanName}})
	{{.CleanName | lowercase}}repo.count = 1
	return {{.CleanName | lowercase}}repo
}


// Get queries a {{.CleanName | lowercase}} by id, throws an error when id is not found
func (repo {{.CleanName}}Repo) Get(id uint64) (*model.{{.CleanName}}, error) {
	value, ok := {{.CleanName | lowercase}}repo.data[id]
	if !ok {
		return nil, fmt.Errorf("get project with id %d, record not found", id)
	}
	return &value, nil
}

// GetAll returns all records ordered by the fields  with isLabel=true
func (repo {{.CleanName}}Repo) GetAll() (model.{{.CleanName}}List, error) {
	var list model.{{.CleanName}}List
	for _,value:=range {{.CleanName | lowercase}}repo.data {
//		{{- range .Fields}}{{if eq .Type "Lookup"}}
//		if {{.Object | lowercase}},_:= {{.Object | lowercase}}repo.Get(value.{{.CleanName}}); {{.Object | lowercase}}!=nil {
//			value.{{.Object}} = {{.Object | lowercase}}.Label()
//		}
//		{{- end}}
		{{if eq .Type "Child"}}
		if {{.Object | lowercase}},_:= {{.Object | lowercase}}repo.Get(value.{{.CleanName}}); {{.Object | lowercase}}!=nil {
			value.{{.Object}} = {{.Object | lowercase}}.Label()
		}
		{{- end}}{{end}}
		list = append(list,value)
	}
			
	return list, nil
}

// Delete deletes the {{.CleanName | lowercase}} with id, throws an error when id is not found
func (repo {{.CleanName}}Repo) Delete(id uint64) error {
	_, ok := {{.CleanName | lowercase}}repo.data[id]
	if !ok {
		return fmt.Errorf("delete project with id '%d', record not found", id)
	}

	delete({{.CleanName | lowercase}}repo.data, id)
	return nil
}

// Update updates all fields in the database table with data from *{{.CleanName}})
func (repo {{.CleanName}}Repo) Update({{.CleanName | lowercase}} *model.{{.CleanName}}) error {
	_, ok := {{.CleanName | lowercase}}repo.data[{{.CleanName | lowercase}}.ID]
	if !ok {
		return fmt.Errorf("update project with id '%d', record not found", {{.CleanName | lowercase}}.ID)
	}	
	{{.CleanName | lowercase}}repo.data[{{.CleanName | lowercase}}.ID] = *{{.CleanName | lowercase}}
	return nil
}

// Insert inserts a new record in the database table with data from *{{.CleanName}})
func (repo {{.CleanName}}Repo) Insert({{.CleanName | lowercase}} *model.{{.CleanName}}) error {
	{{.CleanName | lowercase}}repo.count++
	{{.CleanName | lowercase}}.ID ={{.CleanName | lowercase}}repo.count
	{{.CleanName | lowercase}}repo.data[{{.CleanName | lowercase}}repo.count] = *{{.CleanName | lowercase}}
	return nil
}

// GetLabels returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.CleanName}}Repo) GetLabels() (model.Labels, error) {
	labels := make(model.Labels)
	for _, value := range {{.CleanName | lowercase}}repo.data {
		labels[value.ID] = value.Label()
	}
	return labels, nil
}

{{$name:=.CleanName}}
{{- range .Fields}}{{if eq .Type "Parent"}}
// GetAll{{.Name | plural}}ByParentID returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{$name}}Repo) GetAll{{.Name | plural}}ByParentID(parentID uint64) (model.{{.Object}}List)	{
	list := model.{{.Object}}List{}
	{{.CleanName | plural | lowercase}}, err := {{.Object | lowercase}}repo.GetAll()
	if err!=nil {
		return list
	}
	for _, {{.Object | lowercase}} := range {{.CleanName | plural | lowercase}} {
		if {{.Object | lowercase}}.{{$name}}ID == parentID {
			list = append(list, {{.Object | lowercase}})
		}
	}
	return list
}			
{{- end}}{{end}}

{{end}}
{{end}}


