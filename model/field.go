package model

type Field struct {
  Name string `json:"name" form:"field_name" `
  Type string `json:"type" form:"field_type" `
  Optional bool `json:"optional" form:"field_optional" `
  Length string `json:"length" form:"field_length" `
  Size string `json:"size" form:"field_size" `
  MaxLength int `json:"maxlength" form:"field_maxlength" ` // Specifies the maximum number of characters allowed 
  MinLength int `json:"minlength" form:"field_minlength" ` // Specifies the minimum number of characters required
  Placeholder string `json:"placeholder" form:"field_placeholder" `
  Pattern string `json:"pattern" form:"field_pattern" ` // Specifies a regular expression that an <input> element's value is checked against
  DateTime string `json:"datetime" form:"field-datetime" ` // holds the specifier for the date/time format: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
  Max int `json:"max" form:"field-max" `
  Min int `json:"min" form:"field-min" `
  Step int `json:"step" form:"field-step" `//Specifies the interval between legal numbers in an input field
}