package fsb

import (
	"fmt"
	"strings"
)

type Expression struct {
	condition  string
	bracketFlg bool
}

// Eq
// 等号条件
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

// AND
// AND条件設定
func (e *Expression) AND(exp *Expression) *Expression {
	if strings.Contains(exp.condition, "OR") {
		e.condition = fmt.Sprintf("%s AND (%s)", e.condition, exp.condition)
	} else {
		e.condition = fmt.Sprintf("%s AND %s", e.condition, exp.condition)
	}

	return e
}

// OR
// OR条件設定
func (e *Expression) OR(exp *Expression) *Expression {
	if strings.Contains(exp.condition, "AND") {
		e.condition = fmt.Sprintf("%s OR (%s)", e.condition, exp.condition)
	} else {
		e.condition = fmt.Sprintf("%s OR %s", e.condition, exp.condition)
	}

	e.bracketFlg = true

	return e
}
