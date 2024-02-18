package fsb

import (
	"errors"
	"fmt"
	"strings"
)

type DeleteContainer struct {
	table *TableContainer
	where *Expression
	errs  []error
}

// Delete is a function that initializes a new DeleteContainer instance.
// It takes a pointer to a TableContainer and returns a pointer to a DeleteContainer.
func Delete(table *TableContainer) *DeleteContainer {
	return &DeleteContainer{
		table: table,
	}
}

// Where sets the conditions for the deletion operation.
// It takes a pointer to an Expression and returns a pointer to the DeleteContainer instance.
func (d *DeleteContainer) Where(conditions *Expression) *DeleteContainer {
	d.where = conditions

	return d
}

// ToSQL returns the SQL string representation of the delete operation.
// It checks if there are any errors in the DeleteContainer instance and returns an empty string and the joined errors if any.
// It constructs the SQL elements for the delete operation: "DELETE FROM" and optionally the table name or table alias.
// If a table name is present in DeleteContainer, it appends it to the SQL elements,
// and if a table alias is different from the table name, it appends both the table alias, "AS", and
func (d *DeleteContainer) ToSQL() (string, error) {
	if len(d.errs) > 0 {
		return "", errors.Join(d.errs...)
	}

	sqlElements := []string{"DELETE FROM"}

	if d.table != nil {
		if d.table.name != d.table.bName {
			sqlElements = append(sqlElements, d.table.bName, "AS", d.table.name)
		} else {
			sqlElements = append(sqlElements, d.table.name)
		}
	}

	if d.where != nil {
		sqlElements = append(sqlElements, "WHERE", d.where.condition)
	}

	return fmt.Sprintf("%s;", strings.Join(sqlElements, " ")), nil
}
