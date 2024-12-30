package model

// Relation models the relation between 2 Entities
type Relation struct {
	Name        string `yaml:"name" json:"name" form:"relation-name"`
	Type        string `yaml:"type" json:"type" form:"relation-type"` // 'One-to-Many' | 'Many-to-Many'
	Source      string `yaml:"source" json:"source" form:"relation-source"`
	Destination string `yaml:"destination" json:"destination" form:"relation-destination"`
}

// ContainsEntity indicates if entity 'name' is in that relation
func (r Relation) ContainsEntity(name string) bool {
	if (r.Source == name) || (r.Destination == name) {
		return true
	}
	return false
}
