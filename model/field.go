package model

type Field struct {
	Name        string `yaml:"name" json:"name" form:"field_name" `
	Type        string `yaml:"type" json:"type" form:"field_type" `
	Optional    bool   `yaml:"optional" json:"optional" form:"field_optional" `
	Label       bool   `yaml:"label" json:"label" form:"field_label" `
	Auto        bool   `yaml:"auto" json:"auto" form:"field_auto" `
	Length      string `yaml:"length,omitempty" json:"length" form:"field_length" `
	Size        string `yaml:"size,omitempty" json:"size" form:"field_size" `
	MaxLength   string `yaml:"maxlength,omitempty" json:"maxlength" form:"field_maxlength" ` // Specifies the maximum number of characters allowed
	MinLength   string `yaml:"minlength,omitempty" json:"minlength" form:"field_minlength" ` // Specifies the minimum number of characters required
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder" form:"field_placeholder" `
	Pattern     string `yaml:"pattern,omitempty" json:"pattern" form:"field_pattern" `    // Specifies a regular expression that an <input> element's value is checked against
	DateTime    string `yaml:"datetime,omitempty" json:"datetime" form:"field_datetime" ` // holds the specifier for the date/time format: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
	Max         string `yaml:"max,omitempty" json:"max" form:"field_max" `
	Min         string `yaml:"min,omitempty" json:"min" form:"field_min" `
	Step        string `yaml:"step,omitempty" json:"step" form:"field_step" `       //Specifies the interval between legal numbers in an input field
	Lookup      string `yaml:"lookup,omitempty" json:"lookup" form:"field_lookup" ` // Name of the Lookup list
}
