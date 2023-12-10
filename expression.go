package fsb

import (
	"fmt"
	"strings"
)

// Expression
// AND combines the current expression with another expression using the AND operator.
// If the other expression contains the OR operator, the current expression will be enclosed in brackets.
type Expression struct {
	condition  string
	bracketFlg bool
}

// Eq is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Eq(target string, comp interface{}) *Expression {
	var cond string
	switch v := comp.(type) {
	case string:
		cond = fmt.Sprintf("%s = '%s'", target, v)
	case int:
		cond = fmt.Sprintf("%s = %d", target, v)
	}

	return &Expression{
		condition: cond,
	}
}

// Neq is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately with the "!=" operator.
// The function returns a pointer to an Expression struct initialized with the condition.
func Neq(target string, comp interface{}) *Expression {
	var cond string
	switch v := comp.(type) {
	case string:
		cond = fmt.Sprintf("%s != '%s'", target, v)
	case int:
		cond = fmt.Sprintf("%s != %d", target, v)
	}

	return &Expression{
		condition: cond,
	}
}

// Gt is a function that creates an Expression with a specific condition based on the target and comparison value, where the target is greater than the comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Gt(target string, comp interface{}) *Expression {
	var cond string
	switch v := comp.(type) {
	case string:
		cond = fmt.Sprintf("%s > '%s'", target, v)
	case int:
		cond = fmt.Sprintf("%s > %d", target, v)
	}

	return &Expression{
		condition: cond,
	}
}

// Gte is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately,
// with ">=" as the comparison operator.
// The function returns a pointer to an Expression struct initialized with the condition.
func Gte(target string, comp interface{}) *Expression {
	var cond string
	switch v := comp.(type) {
	case string:
		cond = fmt.Sprintf("%s >= '%s'", target, v)
	case int:
		cond = fmt.Sprintf("%s >= %d", target, v)
	}

	return &Expression{
		condition: cond,
	}
}

// AND sets the condition of the Expression object with the logical AND operator.
// If the exp.condition contains the string "OR", it appends the condition in brackets with the AND operator.
// Otherwise, it appends the condition with the AND operator without brackets.
//
// Example usage:
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.AND(&Expression{condition: "id = 1"})
//	// Output: newExp.condition = "name = 'test' AND id = 1"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.AND(&Expression{condition: "id = 1 OR id = 2"})
//	// Output: newExp.condition = "name = 'test' AND (id = 1 OR id = 2)"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.AND(&Expression{condition: "id = 1 AND id = 2"})
//	// Output: newExp.condition = "name = 'test' AND (id = 1 AND id = 2)"
func (e *Expression) AND(exp *Expression) *Expression {
	if strings.Contains(exp.condition, "OR") {
		e.condition = fmt.Sprintf("%s AND (%s)", e.condition, exp.condition)
	} else {
		e.condition = fmt.Sprintf("%s AND %s", e.condition, exp.condition)
	}

	return e
}

// OR sets the condition of the Expression object with the logical OR operator.
// If the exp.condition contains the string "AND", it appends the condition in brackets with the OR operator.
// Otherwise, it appends the condition with the OR operator without brackets.
//
// It also sets the bracketFlg to true, indicating that the condition contains brackets.
//
// Example usage:
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.OR(&Expression{condition: "id = 1"})
//	// Output: newExp.condition = "name = 'test' OR id = 1"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.OR(&Expression{condition: "id = 1 AND id = 2"})
//	// Output: newExp.condition = "name = 'test' OR (id = 1 AND id = 2)"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.OR(&Expression{condition: "id = 1 OR id = 2"})
//	// Output: newExp.condition = "name = 'test' OR (id = 1 OR id = 2)"
func (e *Expression) OR(exp *Expression) *Expression {
	if strings.Contains(exp.condition, "AND") {
		e.condition = fmt.Sprintf("%s OR (%s)", e.condition, exp.condition)
	} else {
		e.condition = fmt.Sprintf("%s OR %s", e.condition, exp.condition)
	}

	e.bracketFlg = true

	return e
}
