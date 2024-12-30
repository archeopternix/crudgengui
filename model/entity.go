package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
)

type Entity struct {
	Name   string           `yaml:"name" form:"entity-name" `
	Type   string           `yaml:"type" form:"entity-type" ` // 'Entity' || 'Key-Values'
	Fields map[string]Field `yaml:"fields" form:"entity-fields" `
}

func NewEntity() *Entity {
	e := new(Entity)
	e.Fields = make(map[string]Field)
	return e
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
