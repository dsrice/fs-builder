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
	field  []string
	table  *TableContainer
	joins  []*JoinContainer
	where  *Expression
	orders []*OrderContainer
	errs   []error
}

type JoinContainer struct {
	joinType   int
	table      *TableContainer
	conditions []*Expression
}

type OrderContainer struct {
	orderType      int
	orderColumnStr string
}

const (
	inner = 1
	left  = 2
	right = 3
	full  = 4
	cross = 5

	asc  = 1
	desc = 2
)

// Select
// It initializes a new SelectContainer structure.
// It takes all columns that should be selected. If no columns are passed, it assumes '*' (All columns).
func Select(fields ...interface{}) *SelectContainer {
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

	return &SelectContainer{
		table:  &TableContainer{},
		field:  f,
		joins:  []*JoinContainer{},
		orders: []*OrderContainer{},
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

// InnerJoin
// It adds a INNER JOIN to the select statement.
// Parameters:
// - table: the table to join with.
// - conditions...: optional conditions for the join.
// Returns:
// - *SelectContainer: the modified SelectContainer instance with the new join added.
// Usage example:
// sb := fsb.Select().
//
//	From(fsb.Table("users")).
//	InnerJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))
//
// sql, err := sb.ToSQL()
// Output:
// "SELECT * FROM users INNER JOIN tokens ON users.id = '1';"
func (s *SelectContainer) InnerJoin(table *TableContainer, conditions ...*Expression) *SelectContainer {
	join := JoinContainer{
		joinType:   inner,
		table:      table,
		conditions: conditions,
	}

	s.joins = append(s.joins, &join)
	return s
}

func (s *SelectContainer) LeftJoin(table *TableContainer, conditions ...*Expression) *SelectContainer {
	join := JoinContainer{
		joinType:   left,
		table:      table,
		conditions: conditions,
	}

	s.joins = append(s.joins, &join)
	return s
}

func (s *SelectContainer) RightJoin(table *TableContainer, conditions ...*Expression) *SelectContainer {
	join := JoinContainer{
		joinType:   right,
		table:      table,
		conditions: conditions,
	}

	s.joins = append(s.joins, &join)
	return s
}

func (s *SelectContainer) FullJoin(table *TableContainer, conditions ...*Expression) *SelectContainer {
	join := JoinContainer{
		joinType:   full,
		table:      table,
		conditions: conditions,
	}

	s.joins = append(s.joins, &join)
	return s
}

func (s *SelectContainer) CrossJoin(table *TableContainer) *SelectContainer {
	join := JoinContainer{
		joinType: cross,
		table:    table,
	}

	s.joins = append(s.joins, &join)
	return s
}

func (s *SelectContainer) Order(conditions ...interface{}) *SelectContainer {
	orderStr := ""
	for i, condition := range conditions {
		if i > 0 {
			orderStr = fmt.Sprintf("%s, %s", orderStr, ConvertColumn(condition, true))
		} else {
			orderStr = ConvertColumn(condition, true)
		}
	}

	order := OrderContainer{
		orderType:      asc,
		orderColumnStr: orderStr,
	}

	s.orders = append(s.orders, &order)
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
		if s.table.name != s.table.bName {
			sqlElements = append(sqlElements, "FROM", s.table.bName, "AS", s.table.name)
		} else {
			sqlElements = append(sqlElements, "FROM", s.table.name)
		}
	}

	if len(s.joins) > 0 {
		sqlElements = s.createJoinSQL(sqlElements)
	}

	if s.where != nil {
		sqlElements = append(sqlElements, "WHERE", s.where.condition)
	}

	if len(s.orders) > 0 {
		sqlElements = s.createOrderSQL(sqlElements)
	}

	return fmt.Sprintf("%s;", strings.Join(sqlElements, " ")), nil
}

func (s *SelectContainer) createJoinSQL(sqlElements []string) []string {
	for _, join := range s.joins {
		joinTypeStr := ""
		switch join.joinType {
		case inner:
			joinTypeStr = "INNER JOIN"
		case left:
			joinTypeStr = "LEFT JOIN"
		case right:
			joinTypeStr = "RIGHT JOIN"
		case full:
			joinTypeStr = "FULL JOIN"
		case cross:
			joinTypeStr = "CROSS JOIN"
		}

		joinConditions := make([]string, len(join.conditions))
		for i, condition := range join.conditions {
			joinConditions[i] = condition.condition
		}

		tn := join.table.name
		if join.table.name != join.table.bName {
			tn = fmt.Sprintf("%s AS %s", join.table.bName, join.table.name)
		}

		if len(joinConditions) > 0 {
			sqlElements = append(
				sqlElements,
				fmt.Sprintf("%s %s ON %s", joinTypeStr, tn, strings.Join(joinConditions, " AND ")),
			)
		} else {
			sqlElements = append(
				sqlElements,
				fmt.Sprintf("%s %s", joinTypeStr, tn),
			)
		}
	}

	return sqlElements
}

func (s *SelectContainer) createOrderSQL(elements []string) []string {
	orderStr := "ORDER BY"
	for i, order := range s.orders {
		if i > 0 {
			orderStr = fmt.Sprintf("%s, ", orderStr)
		}
		orderStr = fmt.Sprintf("%s %s", orderStr, order.orderColumnStr)

		switch order.orderType {
		case asc:
			orderStr = fmt.Sprintf("%s %s", orderStr, "ASC")
		case desc:
			orderStr = fmt.Sprintf("%s %s", orderStr, "DESC")
		}

		elements = append(elements, orderStr)
	}

	return elements
}
