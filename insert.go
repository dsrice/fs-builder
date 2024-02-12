package fsb

import (
	"errors"
	"fmt"
	"strings"
)

type InsertContainer struct {
	fields []string
	table  *TableContainer
	values []string
	errs   []error
}

// Insert is a function that returns an instance of InsertContainer, which is used to build SQL insert statements.
func Insert(fields ...interface{}) *InsertContainer {
	var f []string

	if len(fields) > 0 {
		for _, l := range fields {
			switch v := l.(type) {
			case string:
				f = append(f, v)
			case *ColumnContainer:
				f = append(f, fmt.Sprintf("%s.%s", v.tName, v.col))
			}
		}
	}

	return &InsertContainer{
		fields: f,
		values: []string{},
	}
}

// Into is a method of InsertContainer that sets the table for the SQL insert statement.
// It takes a *TableContainer as the input parameter and assigns it to the 'table' field of the InsertContainer instance.
// The method returns the modified InsertContainer instance.
func (ic *InsertContainer) Into(table *TableContainer) *InsertContainer {
	ic.table = table

	return ic
}

// Value is a method of InsertContainer that appends a value string to the values slice.
// It takes a variadic number of fields of different types as input parameters.
// The method converts each field to a string representation based on its type and appends it to the value string.
// The value string is then appended to the values slice of the InsertContainer instance.
// The method returns the modified InsertContainer instance.
func (ic *InsertContainer) Value(fields ...interface{}) *InsertContainer {
	values := []string{}
	for _, field := range fields {
		switch t := field.(type) {
		case string:
			values = append(values, fmt.Sprintf("'%s'", t))
		case int:
			values = append(values, fmt.Sprintf("%d", t))
		case float32, float64:
			values = append(values, fmt.Sprintf("%f", t))
		case bool:
			values = append(values, fmt.Sprintf("%v", t))
		default:
			values = append(values, fmt.Sprintf("%v", t))
		}
	}

	ic.values = append(ic.values, strings.Join(values, ", "))

	return ic
}

// ToSQL is a method of InsertContainer that generates a SQL insert statement.
// It returns a string representation of the generated SQL statement and an error, if any.
// Before generating the SQL statement, it checks if there are any errors present in the InsertContainer instance.
// If errors are found, it returns an empty string and joins the errors using the errors.Join function.
func (ic *InsertContainer) ToSQL() (string, error) {
	if len(ic.errs) > 0 {
		return "", errors.Join(ic.errs...)
	}

	sqlElements := []string{"INSERT"}

	if ic.table != nil {
		if ic.table.name != ic.table.bName {
			sqlElements = append(sqlElements, "INTO", ic.table.bName, "AS", ic.table.name)
		} else {
			sqlElements = append(sqlElements, "INTO", ic.table.name)
		}
	}

	if len(ic.fields) > 0 {
		sqlElements = append(sqlElements, "(", strings.Join(ic.fields, ", "), ")")
	}

	if len(ic.values) > 0 {
		sqlElements = append(sqlElements, "VALUES", "(", strings.Join(ic.values, ", "), ")")
	} else {
		ic.errs = append(ic.errs, errors.New("no values provided for insertion"))
	}

	return fmt.Sprintf("%s;", strings.Join(sqlElements, " ")), nil
}
