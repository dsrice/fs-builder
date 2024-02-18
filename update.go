package fsb

import (
	"fmt"
	"strings"
)

type UpdateContainer struct {
	fields map[string]interface{}
	table  *TableContainer
	where  *Expression
	errs   []error
}

// Update is a function that creates a new UpdateContainer object.
// It takes a pointer to a TableContainer as input and returns a pointer to an UpdateContainer.
// The UpdateContainer object is initialized with the provided TableContainer and an empty "fields" map.
// The fields map will be used to store the fields and their respective values to be updated.
func Update(table *TableContainer) *UpdateContainer {
	return &UpdateContainer{
		table:  table,
		fields: make(map[string]interface{}),
	}
}

// Set is a method of UpdateContainer that sets the value of a column in the fields map.
// It takes two parameters, column and value, and returns a pointer to the UpdateContainer.
// The column parameter can either be a string or a *ColumnContainer.
func (u *UpdateContainer) Set(column, value interface{}) *UpdateContainer {
	var c string

	switch v := column.(type) {
	case string:
		c = v
	case *ColumnContainer:
		c = fmt.Sprintf("%s.%s", v.tName, v.col)
	}

	u.fields[c] = value

	return u
}

// SetMap is a method of UpdateContainer that sets the fields map to the given vmap.
// It takes a single parameter, vmap, which is a map[string]interface{}.
// It returns a pointer to the UpdateContainer.
func (u *UpdateContainer) SetMap(vmap map[string]interface{}) *UpdateContainer {
	u.fields = vmap

	return u
}

// Where
// It sets the WHERE clause of the SQL SELECT statement.
func (u *UpdateContainer) Where(conditions *Expression) *UpdateContainer {
	u.where = conditions

	return u
}

func (u *UpdateContainer) ToSQL() (string, error) {
	var setValues []string
	var sql string

	for column, value := range u.fields {
		setValues = append(setValues, fmt.Sprintf("%s = %s", column, ConvertColumn(value, false)))
	}

	if u.table != nil {
		if u.table.name != u.table.bName {
			sql = fmt.Sprintf("UPDATE %s AS %s SET ", u.table.bName, u.table.name)
		} else {
			sql = fmt.Sprintf("UPDATE %s SET ", u.table.name)
		}
	} else {
		u.errs = append(u.errs, fmt.Errorf("no set Table"))
		return "", fmt.Errorf("no set Table")
	}

	sql = fmt.Sprintf("%s %s", sql, strings.Join(setValues, ", "))

	if u.where != nil {
		sql += fmt.Sprintf(" WHERE %s", u.where.condition)
	}

	return sql, nil
}