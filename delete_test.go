package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeleteSuite struct {
	suite.Suite
}

func (s *DeleteSuite) Test_Delete() {
	sb := fsb.Delete(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "DELETE FROM users;", sql)
	assert.Nil(s.T(), err)
}

func (s *DeleteSuite) Test_DeleteWhere() {
	sb := fsb.Delete(fsb.Table("users")).Where(fsb.Eq("id", 1))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "DELETE FROM users WHERE id = 1;", sql)
	assert.Nil(s.T(), err)
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}