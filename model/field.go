package model

import (
	"crudgengui/pkg"
)

type Field struct {
	Name        string `yaml:"name" json:"name" form:"field_name" `                                // Name of the field - see technical ID
	Type        string `yaml:"type" json:"type" form:"field_type" `                                // Type of field
	Id          string `yaml:"id" json:"name" form:"field_id" `                                    // technical ID automatically calculated from the Name
	Required    bool   `yaml:"required" json:"required" form:"field_required" `                    // must be populated
	Label       bool   `yaml:"label" json:"label" form:"field_label" `                             // used as Label for 1..n relations
	Auto        bool   `yaml:"auto" json:"auto" form:"field_auto" `                                // auto field, must not be shown in frontend
	Height      string `yaml:"height,omitempty" json:"height" form:"field_height" `                // Field height for longtext fields
	Size        string `yaml:"size,omitempty" json:"size" form:"field_size" `                      // Size in characters for the field
	MaxLength   string `yaml:"maxlength,omitempty" json:"maxlength" form:"field_maxlength" `       // Specifies the maximum number of characters allowed
	MinLength   string `yaml:"minlength,omitempty" json:"minlength" form:"field_minlength" `       // Specifies the minimum number of characters required
	Placeholder string `yaml:"placeholder,omitempty" json:"placeholder" form:"field_placeholder" ` // Placeholder for entry field
	Pattern     string `yaml:"pattern,omitempty" json:"pattern" form:"field_pattern" `             // Specifies a regular expression that an <input> element's value is checked against
	DateTime    string `yaml:"datetime,omitempty" json:"datetime" form:"field_datetime" `          // holds the specifier for the date/time format: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
	Max         string `yaml:"max,omitempty" json:"max" form:"field_max" `                         // Maximum number
	Min         string `yaml:"min,omitempty" json:"min" form:"field_min" `                         // Minimum number
	Step        string `yaml:"step,omitempty" json:"step" form:"field_step" `                      //Specifies the interval between allowed numbers in an input field
	Decimals    string `yaml:"decimals,omitempty" json:"decimals" form:"field_decimal `            // amount of decimals for number type
	Lookup      string `yaml:"lookup,omitempty" json:"lookup" form:"field_lookup" `                // Name of the Lookup list
	Object      string `yaml:"object,omitempty" json:"object" `                                    // Holding the name of the related object (parent / child)
}

// CleanName removes all non-numeric and non-alphanumeric characters from the input string.
func (f Field) CleanName() string {
	return pkg.CleanID(f.Name)
}
