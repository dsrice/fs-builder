package fsb

import (
	"errors"
	"fmt"
	"strings"
)

// SelectContainer
// This structure represents a SQL SELECT statement.
// It contains fields for the columns being selected (field),
// hose from (table), condition (where), and a list of errors (errs).
type SelectContainer struct {
	field []string
	table *TableContainer
	where *Expression
	errs  []error
}

// Select
// It initializes a new SelectContainer structure.
// It takes all columns that should be selected. If no columns are passed, it assumes '*' (All columns).
func Select(fields ...string) *SelectContainer {
	var f []string

	if len(fields) > 0 {
		f = fields
	}

	return &SelectContainer{
		table: &TableContainer{},
		field: f,
	}
}

// From
// It sets the table from which data has to be selected.
// This method uses a fluent pattern,
// meaning it returns the instance of the container itself,
// allowing the calling of multiple methods in a single line (chaining of function calls).
func (s *SelectContainer) From(table *TableContainer) *SelectContainer {
	s.table = table

	return s
}

// Where
// It sets the WHERE clause of the SQL SELECT statement.
func (s *SelectContainer) Where(conditions *Expression) *SelectContainer {
	s.where = conditions

	return s
}

// ToSQL
// It generates a SQL SELECT statement from the configured SelectContainer structure.
// If any errors exist inside the errs field,
// it will return an empty string and the error.
// The SQL string is composed by appending different components of the select statement.
func (s *SelectContainer) ToSQL() (string, error) {
	if len(s.errs) > 0 {
		return "", errors.Join(s.errs...)
	}

	sqlElements := []string{"SELECT"}

	if len(s.field) > 0 {
		sqlElements = append(sqlElements, strings.Join(s.field, ", "))
	} else {
		sqlElements = append(sqlElements, "*")
	}

	if s.table != nil {
		sqlElements = append(sqlElements, "FROM", s.table.name)
	}

	if s.where != nil {
		sqlElements = append(sqlElements, "WHERE", s.where.condition)
	}

	return fmt.Sprintf("%s;", strings.Join(sqlElements, " ")), nil
}
