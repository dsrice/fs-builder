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

// Test_SelectString tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString() {
	sb := fsb.Select().From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_Field tests the SelectString method in the SelectSuite struct.
// It selects the "id" field from the "users" table and generates a SQL query.
// The generated SQL query should be "SELECT id FROM users;".
// It asserts that the generated SQL query and error are as expected.
func (s *SelectSuite) Test_SelectString_Field() {
	sb := fsb.Select("id").From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT id FROM users;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereInt tests the SelectString_WhereInt method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereInt() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("id", 1))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereString tests the SelectString_WhereString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereString() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test"))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test';", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereAND tests the WhereAND method in the SelectSuite struct.
// It checks if the generated SQL query from SelectString_WhereAnd method matches the expected SQL query.
func (s *SelectSuite) Test_SelectString_WhereAND() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").AND(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' AND id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereOR tests the WhereOR method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereOR() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").OR(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' OR id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereAndOr tests the SelectString method in the SelectSuite struct.
// It creates a SelectContainer object and calls the From method to set the table to "users".
// Then it calls the Where method to set the conditions to "name = 'test' AND (id = 1 OR id = 2)".
// Finally, it calls the ToSQL method to generate the SQL string and checks if it matches the expected value.
// The test also checks if there are no errors returned.
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

// Test_SelectString_WhereOrAnd tests the WhereOrAnd method in the SelectSuite struct.
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

// Test_SelectString_WhereMulti tests the SelectString_WhereMulti method in the SelectSuite struct.
// Selects from the "users" table with the following conditions:
// - "name" equals "test" OR ("id" equals 1 AND "id" equals 2)
// - AND ("login_id" equals "test2" AND "name" equals "name2")
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

// Test_SelectString_InnerJoin tests the SelectString method in the SelectSuite struct.
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

// Test_SelectString_InnerJoinCol tests the SelectString method with InnerJoinCol in the SelectSuite struct.
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

// Test_SelectString_InnerJoinTable tests the SelectString method in the SelectSuite struct
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

// Test_SelectString_LeftJoin tests the LeftJoin method in the SelectContainer struct.
// It joins the "users" table with the "tokens" table using the condition "users.id = '1'".
// It then calls the ToSQL method to generate the SQL query string and checks if it matches the expected query.
func (s *SelectSuite) Test_SelectString_LeftJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		LeftJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users LEFT JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_LeftJoinCol tests the SelectString method in the SelectSuite struct
// with a LeftJoin operation on two tables: "users" and "tokens" based on the condition "users.id
func (s *SelectSuite) Test_SelectString_LeftJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		LeftJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users LEFT JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_LeftJoinTable tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_LeftJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		LeftJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u LEFT JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_RightJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		RightJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users RIGHT JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_RightJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		RightJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users RIGHT JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_RightJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		RightJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u RIGHT JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_FullJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		FullJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users FULL JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_FullJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		FullJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users FULL JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_FullJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		FullJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u FULL JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

func TestSelectSuite(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}
