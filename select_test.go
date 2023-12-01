package fsb_test

import (
	"fsb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type user struct {
	ID int
}

func (u *user) TableName() string {
	return "user"
}

type SelectSuite struct {
	suite.Suite
}

func (s *SelectSuite) Test_Select_Model() {
	sb := fsb.Select().From(&user{})
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM user;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_Select_String() {
	sb := fsb.Select().From("user")
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM user;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_Select_Failed() {
	sb := fsb.Select().From(1)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "", sql)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), "failed set table", err.Error())
}

func TestSelectSuite(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}
