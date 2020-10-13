package client

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ConfigTest struct{}

var _ = Suite(&ConfigTest{})

func (s *ConfigTest) SetUpSuite(c *C) {
}

func (s *ConfigTest) TearDownSuite(c *C) {
}

func (s *ConfigTest) TestGenUUID(c *C) {
	c.Assert(generateUUID(), Not(Equals), generateUUID())
}

func (s *ConfigTest) TestFileExist(c *C) {
	c.Check(fileExist("go.mod") || fileExist("config.go"), Equals, true)
	c.Check(fileExist("fileNotExist"), Equals, false)
}
