{{define "databaserepotesttypes"}}
{{- if eq .Kind "Text"}}{{.Name}}: getText(15),{{end}}
{{- if eq .Kind "Password"}}{{.Name}}: getText(15),{{end}}
{{- if eq .Kind "Integer"}}{{- if ne .Name "ID"}}{{.Name}}: rand.Uint64(),{{end}}{{end}}
{{- if eq .Kind "Number"}}{{.Name}}: rand.Float32(),{{end}}
{{- if eq .Kind "Boolean"}}{{.Name}}:true,{{end}}
{{- if eq .Kind "Email"}}{{.Name}}: getEmail(),{{end}}
{{- if eq .Kind "Tel"}}	{{.Name}}: getText(12), {{end}}
{{- if eq .Kind "Longtext"}}{{.Name}}: getText(50),{{end}}
{{- if eq .Kind "Time"}}{{.Name}}:time.Now(),{{end}}
{{- if eq .Kind "Lookup"}}{{.Name}}: 1,{{end}}
{{- if eq .Kind "Child"}}{{.Name}}:1,{{end}}
{{- end}}