// Package model holds all entities which are needed for generation of the
// target application
package model

import (
	"time"
)

// Entity relates to an 'Object' or struct
type Entity struct {
	Name   string           `yaml:"name"`
	Fields map[string]Field `yaml:"fields"`
	Kind   string           `yaml:"type,omitempty"` // 0..Normal, 1..Lookup 2..Many2Many
}

func NewEntity() Entity {
	e := Entity{}
	e.Fields = make(map[string]Field)
	return e
}

// AddField add fields to an entity
func (e *Entity) AddField(f Field) {
	e.Fields[f.Name] = f
}

// TimeStamp needed for file generation. Will be added in the header of each file
// to track the creation date and time of each file
func (e Entity) TimeStamp() string {
	return time.Now().Format("01.02.2006 15:04:05")
}
