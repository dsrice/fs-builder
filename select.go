package fsb

import (
	"errors"
	"fmt"
	"strings"
)

type SelectContainer struct {
	field []string
	table *TableContainer
	where *Expression
	errs  []error
}

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

func (s *SelectContainer) From(table *TableContainer) *SelectContainer {
	s.table = table

	return s
}

func (s *SelectContainer) Where(conditions *Expression) *SelectContainer {
	s.where = conditions

	return s
}

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
