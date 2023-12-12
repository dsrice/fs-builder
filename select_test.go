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
	sb := fsb.Select("id").From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT id FROM users;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereInt() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("id", 1))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE id = 1;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereString() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test"))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test';", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereAND() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").AND(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' AND id = 1;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereOR() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").OR(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' OR id = 1;", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereAndOr() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").
			AND(
				fsb.Eq("id", 1).OR(fsb.Eq("id", 2)),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' AND (id = 1 OR id = 2);", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereOrAnd() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").
			OR(
				fsb.Eq("id", 1).AND(fsb.Eq("id", 2)),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' OR (id = 1 AND id = 2);", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_WhereMulti() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").
			OR(
				fsb.Eq("id", 1).AND(fsb.Eq("id", 2)),
			).
			AND(
				fsb.Eq("login_id", "test2").AND(fsb.Eq("name", "name2")),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users WHERE name = 'test' OR (id = 1 AND id = 2) AND login_id = 'test2' AND name = 'name2';",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_InnerJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		InnerJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users INNER JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_InnerJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		InnerJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users INNER JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_InnerJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		InnerJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u INNER JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func TestSelectSuite(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}
