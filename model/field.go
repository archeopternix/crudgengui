package model

type Field struct {
  Name string `json:"name" form:"field-name" `
  Type string `json:"type" form:"field-type" `
  Optional bool `json:"optional" form:"field-optional" `
  Length int `json:"length" form:"field-length" `
  Size int `json:"size" form:"field-size" `
}

type TextField struct {
  MaxLength int `json:"maxlength" form:"field-maxlength" ` // Specifies the maximum number of characters allowed 
  MinLength int `json:"minlength" form:"field-minlength" ` // Specifies the minimum number of characters required
  Size int `json:"size" form:"field-size" ` //Specifies the width, in characters, of an <input> element
  Placeholder string `json:"placeholder" form:"field-placeholder" `
  Pattern string `json:"pattern" form:"field-pattern" ` // Specifies a regular expression that an <input> element's value is checked against
}

type NumberField {
  Max int `json:"max" form:"field-max" `
  Min int `json:"min" form:"field-min" `
  Step int `json:"step" form:"field-step" `//Specifies the interval between legal numbers in an input field
}