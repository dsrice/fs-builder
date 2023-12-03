package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
	"unsafe"
)

type TableSuite struct {
	suite.Suite
}

func (s *TableSuite) Test_SetTable() {
	t := fsb.Table("users")

	v := reflect.ValueOf(t).Elem()
	pstrv := v.FieldByName("name")
	ps := (*string)(unsafe.Pointer(pstrv.UnsafeAddr()))

	assert.Equal(s.T(), "users", *ps)

	pstrv = v.FieldByName("bName")
	ps = (*string)(unsafe.Pointer(pstrv.UnsafeAddr()))

	assert.Equal(s.T(), "users", *ps)

	t = fsb.Table("users").As("t")

	v = reflect.ValueOf(t).Elem()
	pstrv = v.FieldByName("name")
	ps = (*string)(unsafe.Pointer(pstrv.UnsafeAddr()))

	assert.Equal(s.T(), "t", *ps)

	pstrv = v.FieldByName("bName")
	ps = (*string)(unsafe.Pointer(pstrv.UnsafeAddr()))

	assert.Equal(s.T(), "users", *ps)
}

func TestTableSuite(t *testing.T) {
	suite.Run(t, new(TableSuite))
}