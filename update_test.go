package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UpdateSuite struct {
	suite.Suite
}

// Test_Update is a test function that tests the Update method of UpdateSuite.
func (s *UpdateSuite) Test_Update() {
	sb := fsb.Update(fsb.Table("users")).Set("id", 1)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "UPDATE users SET id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_UpdateMulti is a test function that tests the UpdateMulti method of UpdateSuite.
func (s *UpdateSuite) Test_UpdateMulti() {
	sb := fsb.Update(fsb.Table("users")).Set("name", "test").Set("id", 1)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "UPDATE users SET name = 'test', id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_UpdateSetMap is a test function that tests the UpdateSetMap method of UpdateSuite.
func (s *UpdateSuite) Test_UpdateSetMap() {
	vmap := make(map[string]interface{})
	vmap["id"] = 1
	vmap["name"] = "test"

	sb := fsb.Update(fsb.Table("users")).SetMap(vmap)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "UPDATE users SET id = 1, name = 'test';", sql)
	assert.Nil(s.T(), err)
}

// Test_UpdateWhere is a test function that tests the UpdateWhere method of UpdateSuite.
func (s *UpdateSuite) Test_UpdateWhere() {
	sb := fsb.Update(fsb.Table("users")).Set("name", "test").Where(fsb.Eq("id", 1))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "UPDATE users SET name = 'test' WHERE id = 1;", sql)
	assert.Nil(s.T(), err)
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}