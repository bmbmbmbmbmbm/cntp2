package db

import (
	"database/sql"
	"fmt"
)

type SelectQuery interface {
	Select(table string)
	Where(field string, value any)
	And(field string, value any)
	Or(field string, value any)
	Build(*sql.DB) (*sql.Rows, error)
}

type SelectBuilder struct {
	statement string
}

func (sb *SelectBuilder) Select(table string) *SelectBuilder {
	if len(sb.statement) != 0 {
		sb.statement = fmt.Sprintf("FROM %s ", table)
	}
	return sb
}

func (sb *SelectBuilder) Where(field string, value any) *SelectBuilder {
	sb.statement += fmt.Sprintf("WHERE %s=", field)
	switch value.(type) {
	case string:
		sb.statement += fmt.Sprintf("%s", value)
	case int:
		sb.statement += fmt.Sprintf("%d", value)
	case bool:
		sb.statement += fmt.Sprintf("%s", value)
	}
	return sb
}

func (sb SelectBuilder) Build(db *sql.DB) (*sql.Rows, error) {
	return db.Query(sb.statement)
}
