package fsb

import (
	"fmt"
	"fsb/dataset"
)

type SelectContainer struct {
	field []string
	table *dataset.TableContainer
	errs  []error
}

func Select() *SelectContainer {
	return &SelectContainer{}
}

func (s *SelectContainer) From(table interface{}) *SelectContainer {
	tc := dataset.TableContainer{}

	switch table.(type) {
	case dataset.Table:
		p := table.(dataset.Table)
		tc.Name = p.TableName()
	case string:
		tc.Name = table.(string)
	}

	s.table = &tc
	return s
}

func (s *SelectContainer) ToSQL() string {
	sql := "SELECT *"

	if s.table != nil {
		sql = fmt.Sprintf("%s FROM %s", sql, s.table.Name)
	}

	return fmt.Sprintf("%s;", sql)
}
