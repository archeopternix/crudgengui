package model

type Field struct {
  Name string `json:"name" form:"field-name" `
  Type string `json:"type" form:"field-type" `
  Optional bool `json:"optional" form:"field-optional" `
  Length int `json:"length" form:"field-length" `
  Size int `json:"size" form:"field-size" `
}
