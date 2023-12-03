package fsb

import "fmt"

type Expression struct {
	condition string
}

func Eq(target string, comp interface{}) *Expression {
	var cond string
	switch comp.(type) {
	case string:
		cond = fmt.Sprintf("%s = '%s'", target, comp.(string))
	case int:
		cond = fmt.Sprintf("%s = %d", target, comp.(int))
	}

	return &Expression{
		condition: cond,
	}
}