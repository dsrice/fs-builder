package fsb

import (
	"errors"
	"fmt"
	"strings"
)

type TruncateContainer struct {
	table *TableContainer
	errs  []error
}

// Truncate is a function that creates a new TruncateContainer which represents a TRUNCATE TABLE statement in SQL.
// The TruncateContainer contains a reference to a TableContainer representing the table to be truncated.
func Truncate(table *TableContainer) *TruncateContainer {
	return &TruncateContainer{
		table: table,
	}
}

func (t *TruncateContainer) ToSQL() (string, error) {
	if len(t.errs) > 0 {
		return "", errors.Join(t.errs...)
	}

	sqlElements := []string{"TRUNCATE TABLE"}

	if t.table != nil {
		if t.table.name != t.table.bName {
			sqlElements = append(sqlElements, t.table.bName, "AS", t.table.name)
		} else {
			sqlElements = append(sqlElements, t.table.name)
		}
	}

	return fmt.Sprintf("%s;", strings.Join(sqlElements, " ")), nil
}
