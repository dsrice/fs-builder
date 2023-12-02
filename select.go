package fsb

import (
	"errors"
	"fmt"
	"fsb/dataset"
)

type SelectContainer struct {
	field []string
	table *dataset.TableContainer
	errs  []error
}

func Select() *SelectContainer {
	return &SelectContainer{table: &dataset.TableContainer{}}
}

func (s *SelectContainer) From(table interface{}) *SelectContainer {
	switch t := table.(type) {
	case dataset.Table:
		s.table.Name = table.(dataset.Table).TableName()
	case string:
		s.table.Name = t
	default:
		s.errs = []error{fmt.Errorf("failed set table")}
	}

	return s
}

func (s *SelectContainer) ToSQL() (string, error) {
	sql := "SELECT *"

	if len(s.errs) > 0 {
		return "", errors.Join(s.errs...)
	}

	if s.table != nil {
		sql = fmt.Sprintf("%s FROM %s", sql, s.table.Name)
	}

	return fmt.Sprintf("%s;", sql), nil
}