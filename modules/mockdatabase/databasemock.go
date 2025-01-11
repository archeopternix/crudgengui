{{define "databasemock" -}}
// databasemock.go
package database

import (
	"{{.Name}}/model"

)


// GetMockEnv opens database with path or ":memory:" for in memory database
func GetMockEnv() *model.Env {
	env := &model.Env{
	{{range .Entities}}		
		{{.CleanName | plural}}:    New{{.CleanName}}Repo(),
	{{- end}}
	}
	return env
}

{{end}}