// Package model holds all entities which are needed for generation of the
// target application
package model

// Relation holds the definition for parent - child relationships.
// When parsed by Application additional fields will be added to the child and parent
// entities
type Relation struct {
	Name   string `yaml:"name" json:"name"`
	Parent string `yaml:"source" json:"source"`
	Child  string `yaml:"destination" json:"destination"`
	Kind   string `yaml:"type" json:"type"` // "one2many", "many2many"
}
