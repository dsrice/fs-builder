package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TruncateSuite struct {
	suite.Suite
}

func (s *TruncateSuite) Test_Truncate() {
	sb := fsb.Truncate(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "TRUNCATE TABLE users;", sql)
	assert.Nil(s.T(), err)
}

func TestTruncateSuite(t *testing.T) {
	suite.Run(t, new(TruncateSuite))
}
