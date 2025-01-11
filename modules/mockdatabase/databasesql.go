{{define "databasesql" -}}
// databasesql.go
package database

import (
	"fmt"
	"log"

	"{{.Name}}/model"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// GetSQLnv opens database with path or ":memory:" for in memory database
func GetSQLEnv(name string) *model.Env {
	// open and connect at the same time, panicing on error
	db = sqlx.MustConnect("sqlite3", name)
	env := &model.Env{
	{{range .Entities}}		
		{{.CleanName | plural}}:    New{{.CleanName}}SQL(db),
	{{- end}}
	}

	// Create the table if it doesn't exist
	schema := `
	{{range .Entities}}
    CREATE TABLE IF NOT EXISTS {{.CleanName | lowercase}} (
    		id INTEGER PRIMARY KEY AUTOINCREMENT
    		{{range .Fields}}{{if ne .CleanName "ID"}}{{.CleanName | lowercase}} {{.GetDatabaseType}}, {{end}}
		{{end}}
		created_at DATETIME
    );
    {{- end}}
    `
	db.MustExec(schema)

	fmt.Println("Database created")
	return env
}

func ShowSchema() {
	var tables = []string{ {{range .Entities}} "{{.CleanName | lowercase}}",{{end}} }
	// Show the schema of a specific table, e.g., "haus"
	for _,tableName := range  tables {
		query := fmt.Sprintf("PRAGMA table_info(%s);", tableName)
		var columns []struct {
			Cid        int     `db:"cid"`
			Name       string  `db:"name"`
			Type       string  `db:"type"`
			NotNull    int     `db:"notnull"`
			DefaultVal *string `db:"dflt_value"`
			PrimaryKey int     `db:"pk"`
		}
	
		if err := db.Select(&columns, query); err != nil {
			log.Fatalln(err)
		}
	
		fmt.Printf("Schema of table %s:\n", tableName)
		for _, col := range columns {
			fmt.Printf("Column: %s, Type: %s, NotNull: %d, Default: %v, PrimaryKey: %d\n",
				col.Name, col.Type, col.NotNull, col.DefaultVal, col.PrimaryKey)
		}	
		}
}
{{end}}