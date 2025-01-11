{{define "sqlrepo" -}}
// Package mockdatabase contains structures and function for mock database access
// Generated code - do not modify it will be overwritten!!
// Time: {{.TimeStamp}}
package database

import (
	"fmt"
	model "{{.AppName}}/model"
	"github.com/jmoiron/sqlx"
)

{{with .Entity}}

// {{.CleanName}}SQL is the interface for a {{.CleanName}} repository that will persist 
// and retrieve data and has to be implemented for concrete Databases 
// (e.g. db *sqlx.DB) or other respositories
type {{.CleanName}}SQL struct{
	db *sqlx.DB
}


// New{{.CleanName}}Repo creates a new instance of {{.CleanName}}Repo and initializes the 
// map and sets counter for ID to 1
func New{{.CleanName}}SQL(db *sqlx.DB) *{{.CleanName}}SQL {
	return &{{.CleanName}}SQL{db: db}
}


// Get queries a {{.CleanName}} by id, returns an error when id is not found
func (repo *{{.CleanName}}SQL) Get(id uint64) (*model.{{.CleanName}}, error) {
	var {{.CleanName | lowercase}} model.{{.CleanName}}
	err := repo.db.Get(&{{.CleanName | lowercase}}, "SELECT * FROM {{.CleanName | lowercase}} WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("get project with id %d, record not found: %w", id, err)
	}
	return &{{.CleanName | lowercase}}, nil
}


// GetAll returns all records ordered by the fields with isLabel=true
func (repo *{{.CleanName}}SQL) GetAll() (model.{{.CleanName}}List, error) {
	var list model.{{.CleanName}}List

	rows, err := db.Queryx("SELECT * FROM {{.CleanName | lowercase}}")
	if err != nil {
		return nil, fmt.Errorf("GetAll, query error: %w", err)
	}
	for rows.Next() {
		var h model.{{.CleanName}}
		if err := rows.StructScan(&h); err != nil {
			return nil, fmt.Errorf("GetAll, bind error: %w", err)
		}
		list = append(list, h)
	}

	return list, nil
}

// Delete deletes the {{.CleanName}} with id, throws an error when id is not found
func (repo *{{.CleanName}}SQL) Delete(id uint64) error {
	result, err := repo.db.Exec("DELETE FROM {{.CleanName | lowercase}} WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete {{.CleanName}} with id '%d', error: %w", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete {{.CleanName}} with id '%d', error: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("delete {{.CleanName}} with id '%d', record not found", id)
	}
	return nil
}

// Update updates all fields in the database table with data from *{{.CleanName}}
func (repo *{{.CleanName}}SQL) Update({{.CleanName | lowercase}} *model.{{.CleanName}}) error {
	update :={{- $name := .CleanName }}"UPDATE {{.CleanName | lowercase| plural}} SET 
{{- range $index, $element := .Fields}}{{if ne .CleanName "ID"}}{{if ne .Type "Parent"}}
{{- if gt $index 0}},{{end}}{{$element.CleanName | lowercase}}= :{{$element.CleanName | lowercase}} {{end}}{{end}} 
{{- end}}WHERE id= :id" 

	// query := `UPDATE {{.CleanName | lowercase}} SET name = :name, email = :emaila, created_at = :created_at, datum = :datum, tuere = :tuere WHERE id = :id`
	_, err := repo.db.NamedExec(update, {{.CleanName | lowercase}})
	if err != nil {
		return fmt.Errorf("update {{.CleanName}} with id '%d', error: %w", {{.CleanName | lowercase}}.ID, err)
	}
	return nil
}


// Insert inserts a new record in the database table with data from *{{.CleanName}}
func (repo *{{.CleanName}}SQL) Insert({{.CleanName | lowercase}} *model.{{.CleanName}}) error {
	insert:="INSERT INTO {{.CleanName | lowercase | plural}} (
{{- range $index, $element := .Fields}}{{if ne $element.CleanName "ID"}}{{if gt $index 0}},{{end}} {{$element.CleanName | lowercase}}{{end}}{{end}}, created_at) VALUES (
{{- range $index, $element := .Fields}}{{if ne $element.CleanName "ID"}}{{if gt $index 0}},{{end}} :{{$element.CleanName | lowercase}}{{end}}{{end}}, :DEFAULT )"


//	query := `INSERT INTO haus (name, email, created, datum, tuere) VALUES (:name, :email, :created, :datum, :tuere)`
	result, err := repo.db.NamedExec(insert, {{.CleanName | lowercase}})
	if err != nil {
		return fmt.Errorf("1 insert {{.CleanName}}, error: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("2 insert {{.CleanName}}, error: %w", err)
	}
	{{.CleanName | lowercase}}.ID = uint64(id)
	return nil
}




// GetLabels returns a map with the key id and the value of
// all fields tagged with isLabel=true and separated by a blank
func (repo {{.CleanName}}SQL) GetLabels() (model.Labels, error) {
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
func (repo {{$name}}SQL) GetAll{{.Name | plural}}ByParentID(parentID uint64) (model.{{.Object}}List)	{
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


