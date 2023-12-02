package fsb

import (
	"errors"
	"fmt"
	"fsb/dataset"
	"reflect"
	"strings"
)

type SelectContainer struct {
	field []string
	table *dataset.TableContainer
	errs  []error
}

func Select() *SelectContainer {
	return &SelectContainer{
		table: &dataset.TableContainer{},
		field: []string{},
	}
}

func (s *SelectContainer) Field(fields ...string) *SelectContainer {
	s.field = fields

	return s
}

func (s *SelectContainer) From(table interface{}) *SelectContainer {
	switch t := table.(type) {
	case dataset.Table:
		d := table.(dataset.Table)
		s.getStructField(d)
		s.table.Dataset = &d
		s.table.Name = d.TableName()
	case string:
		s.table.Name = t
	default:
		s.errs = []error{fmt.Errorf("failed set table")}
	}

	return s
}

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
		sqlElements = append(sqlElements, "FROM")
		sqlElements = append(sqlElements, s.table.Name)
	}

	return fmt.Sprintf("%s;", strings.Join(sqlElements, " ")), nil
}

func (s *SelectContainer) getStructField(d dataset.Table) {
	v := reflect.ValueOf(d)
	i := reflect.Indirect(v)
	it := i.Type()
	nf := it.NumField()

	var resultList []string

	for i := 0; i < nf; i++ {
		field := it.Field(i).Tag.Get("db")
		resultList = append(resultList, field)
	}

	s.field = resultList
}