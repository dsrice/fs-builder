package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	ID int
}

func (u *User) TableName() string {
	return "user"
}

func Test_Select(t *testing.T) {
	sb := fsb.Select().From(&User{})
	sql := sb.ToSQL()

	assert.Equal(t, "SELECT * FROM user;", sql)

}
