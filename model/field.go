package model

type Field struct {
	Name     string `yaml:"name" json:"name" form:"field_name" `
	Type     string `yaml:"type" json:"type" form:"field_type" `
	Required bool   `yaml:"required" json:"required" form:"field_required" `
	Label    bool   `yaml:"label" json:"label" form:"field_label" `
	Auto     bool   `yaml:"auto" json:"auto" form:"field_auto" `
	// Field height for longtext fields
	Height string `yaml:"height,omitempty" json:"height" form:"field_height" `
	// Size in characters for the field
	Size string `yaml:"size,omitempty" json:"size" form:"field_size" `
	// Specifies the maximum number of characters allowed
	MaxLength string `yaml:"maxlength,omitempty" json:"maxlength" form:"field_maxlength" `
	// Specifies the minimum number of characters required
	MinLength   string `yaml:"minlength,omitempty" json:"minlength" form:"field_minlength" `
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder" form:"field_placeholder" `
	// Specifies a regular expression that an <input> element's value is checked against
	Pattern string `yaml:"pattern,omitempty" json:"pattern" form:"field_pattern" `
	// holds the specifier for the date/time format: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
	DateTime string `yaml:"datetime,omitempty" json:"datetime" form:"field_datetime" `
	Max      string `yaml:"max,omitempty" json:"max" form:"field_max" `
	Min      string `yaml:"min,omitempty" json:"min" form:"field_min" `
	//Specifies the interval between allowed numbers in an input field
	Step string `yaml:"step,omitempty" json:"step" form:"field_step" `
	// amount of decimals for number type
	Decimals string `yaml:"decimals,omitempty" json:"decimals" form:"field_decimal `
	// Name of the Lookup list
	Lookup string `yaml:"lookup,omitempty" json:"lookup" form:"field_lookup" `
	// Holding the name of the related object (parent / child)
	Object string `yaml:"object,omitempty" json:"object" `
}
