package model

import (

)

// Relation models the relation between 2 Entities
type Relation struct {
	Name        string `json:"name" form:"relation-name"`
	Type        string `json:"type" form:"relation-type"` // 'One-to-Many' | 'Many-to-Many'
	Source      string `json:"source" form:"relation-source"`
	Destination string `json:"destination" form:"relation-destination"`
}

// ContainsEntity indicates if entity 'name' is in that relation
func (r Relation) ContainsEntity(name string) bool {
  if (r.Source == name) || (r.Destination == name) {
    return true
  }
  return false
}
