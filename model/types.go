package model

import (
  "time"
)
var BaseTypes =[]string{"int","bool","string","time","float"}
var TemplateTypes =[]string{"Frontend","Api","Database","Model", "App"}


type FieldTypeDefinition struct {
  Name string
  Optional bool
  HasLength bool
  DefaultLength int
  HasSize bool
  DefaultSize int
  HasMin bool
  DefaultMin int
  HasMax bool
  DefaultMax int
  HasStep bool
  DefaultStep float32
  HasDigits bool
  DefaultDigits int
  Format string // uses fmt.Printf format and special format for dates e.g. "2006-01-02 15:04:05"
  BaseType string // int, bool, string, time.date, float
  Icon string // e.g. fa-euro-sign shown in entry fields

  Version int
  ChangeDate time.Time
}

/*type FieldTypeDefinitions map[string]FieldTypeDefinition

type FieldTemplates map[string]FieldDefinition

func NewFieldTemplates() *FieldTemplates {
  ft:= make(FieldTemplates) 
  return ft
}
*/
// https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-repository-content
type GitRepo struct {
  GitName string
  GitPath string
  GitSHA string
  GitRawLink string // Download link
}


type BaseTemplate struct {
  Name string
  Type string // TemplateTypes e.g. Frontend...
  Content string
  
  Version int
  ChangeDate time.Time
}

type FieldTemplate struct {
  Name string
  Type string // TemplateTypes e.g. Frontend...
  FieldTypeName string
  Content string
  
  Version int
  ChangeDate time.Time
}