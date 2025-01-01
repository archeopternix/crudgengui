// Package model holds all entities which are needed for generation of the
// target application
package model

// Field is each and every single attribute.
// Object is empty except in case type=lookup or child keeps the name of the Object
type Field struct {
	Name     string `yaml:"name"`
	Kind     string `yaml:"type"` // string, int, bool, lookup, tel, email
	Required bool   `yaml:"required,omitempty"`
	Label    bool   `yaml:"label,omitempty"`  // when true is the shown text for select boxes
	Object   string `yaml:"object,omitempty"` // for lookup, child relations - mappingtable for many2many relations
	Length   int    `yaml:"length,omitempty"`
	Size     int    `yaml:"size,omitempty"` // for textarea size = cols
	Lookup   string `yaml:"lookup,omitempty"`

	Step int `yaml:"step,omitempty"` //for Number fields
	Min  int `yaml:"min,omitempty"`  //for Number fields
	Max  int `yaml:"max,omitempty"`  //for Number fields

	Rows int `yaml:"rows,omitempty"` //for textarea
}
