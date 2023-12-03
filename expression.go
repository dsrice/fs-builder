package fsb

import "fmt"

type Expression struct {
	condition string
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