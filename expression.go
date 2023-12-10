package fsb

import (
	"fmt"
	"strconv"
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
	cond := createCondition(target, comp, "=")

	return &Expression{
		condition: cond,
	}
}

// Neq is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately with the "!=" operator.
// The function returns a pointer to an Expression struct initialized with the condition.
func Neq(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "!=")

	return &Expression{
		condition: cond,
	}

}

// Gt is a function that creates an Expression with a specific condition based on the target and comparison value, where the target is greater than the comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Gt(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, ">")

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
	cond := createCondition(target, comp, ">=")

	return &Expression{
		condition: cond,
	}

}

// Lt is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately, with a less than (<) symbol.
// The function returns a pointer to an Expression struct initialized with the condition.
func Lt(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "<")

	return &Expression{
		condition: cond,
	}
}

// Lte is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Lte(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "<=")

	return &Expression{
		condition: cond,
	}
}

// Like is a function that creates an Expression with a specific condition based on the target and comparison value using the "LIKE" operator.
// It handles only string comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Like(target string, comp string) *Expression {
	cond := createCondition(target, comp, "LIKE")

	return &Expression{
		condition: cond,
	}
}

// Nlike is a function that creates an Expression struct with a specific condition based on the target and comparison value.
// It uses the "NOT LIKE" sign to build the condition.
// The function returns a pointer to the Expression struct initialized with the condition.
func Nlike(target string, comp string) *Expression {
	cond := createCondition(target, comp, "NOT LIKE")

	return &Expression{
		condition: cond,
	}
}

// Sm is a function that creates an Expression with a specific condition based on the target and comparison value.
// It converts the comparison value to a SQL-like pattern using the toSqlLikePattern function.
// The condition is built using the createCondition function with the target, converted comparison value, and "LIKE" sign.
// The function returns a pointer to an Expression struct initialized with the condition.
func Sm(target string, comp interface{}) *Expression {
	cond := createCondition(target, toSqlLikePattern(comp), "LIKE")

	return &Expression{
		condition: cond,
	}
}

// Nsm is a function that creates an Expression with a specific condition based on the target and comparison value.
// It converts the comparison value to a SQL LIKE pattern using the toSqlLikePattern function.
// The condition is built using the createCondition function with the target, SQL LIKE pattern, and "NOT LIKE" sign.
// The function returns a pointer to an Expression struct initialized with the condition.
func Nsm(target string, comp interface{}) *Expression {
	cond := createCondition(target, toSqlLikePattern(comp), "NOT LIKE")

	return &Expression{
		condition: cond,
	}
}

func Bm(target string, comp interface{}) *Expression {
	cond := createCondition(target, toSqlLikePattern(comp), "LIKE")

	return &Expression{
		condition: cond,
	}
}

// createCondition is a function that takes a target string, a comparison value of type string or int, and a sign string.
// It builds and returns a condition string based on the target, comparison value, and sign provided.
// If the comparison value is a string, it formats the condition as "%s %s '%s'", where the target, sign, and comparison value are substituted accordingly.
// If the comparison value is an int, it formats the condition as "%s %s %d", where the target, sign, and comparison value are substituted accordingly.
// If the comparison value is neither a string nor an int, it returns an empty string.
// Example usage:
//
//	cond := createCondition("age", 25, "=") // Returns "age = 25"
//	cond := createCondition("name", "John", "=") // Returns "name = 'John'"
func createCondition(target string, comp interface{}, sign string) string {
	switch v := comp.(type) {
	case string:
		return fmt.Sprintf("%s %s '%s'", target, sign, v)
	case int:
		return fmt.Sprintf("%s %s %d", target, sign, v)
	default:
		return ""
	}
}

// toSqlLikePattern is a function that converts a comparison value to a SQL LIKE pattern.
// It takes an interface{} as an argument and returns a string.
// If the comparison value is a string, it appends '%' at the end.
// If the comparison value is an int, it converts it to a string and appends '%' at the end.
// For any other type, it returns an empty string.
func toSqlLikePattern(comp interface{}) string {
	switch v := comp.(type) {
	case string:
		return v + "%"
	case int:
		return strconv.Itoa(v) + "%"
	default:
		return ""
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
