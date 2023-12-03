package fsb_test

import (
	"fsb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SelectSuite struct {
	suite.Suite
}

func (s *SelectSuite) Test_SelectString() {
	sb := fsb.Select().From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_Field() {
	sb := fsb.Select().Field("id").From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT id FROM users;", sql)
	assert.Nil(s.T(), err)
}
func TestSelectSuite(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}