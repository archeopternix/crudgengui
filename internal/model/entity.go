package model

import (
	"crudgengui/pkg"
	"slices"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
)

type Entity struct {
	Name   string  `yaml:"name" form:"entity-name" `
	Type   string  `yaml:"type" form:"entity-type" ` // 'Entity' || 'Key-Values'
	Fields []Field `yaml:"fields" form:"entity-fields" `
}

func NewEntity() *Entity {
	e := new(Entity)
	return e
}

// CleanName removes all non-numeric and non-alphanumeric characters from the input string.
func (e Entity) CleanName() string {
	return pkg.CleanString(e.Name)
}

// TimeStamp needed for file generation. Will be added in the header of each file
// to track the creation date and time of each file
func (e Entity) TimeStamp() string {
	return time.Now().Format("01.02.2006 15:04:05")
}

func (e *Entity) Add(f Field) {
	e.Fields = append(e.Fields, f)
}

func (e Entity) GetFieldIndexByName(name string) int {
	for i, f := range e.Fields {
		if f.Name == name {
			return i
		}
	}
	return -1
}

func (e Entity) GetFieldByName(name string) *Field {
	f := new(Field)
	if i := e.GetFieldIndexByName(name); i > -1 {
		f = &e.Fields[i]
		return f
	}
	return nil
}

func (e *Entity) DeleteFieldByName(name string) {
	i := e.GetFieldIndexByName(name)
	e.Fields = slices.Delete(e.Fields, i, i+1)
}

type ErrorMap map[string]error

func (e *Entity) IsValid() (bool, ErrorMap) {
	valid := true
	errormap := make(ErrorMap)

	if err := validation.Validate(e.Name, validation.Required, is.Alpha); err != nil {
		valid = false
		errormap["Name"] = err
	}
	if err := validation.Validate(e.Type, validation.Required, is.Alpha); err != nil {
		valid = false
		errormap["Type"] = err
	}

	return valid, errormap
}
