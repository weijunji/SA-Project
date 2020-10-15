package server

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestConfig(t *testing.T) { TestingT(t) }

type ConfigTest struct{}

var _ = Suite(&ConfigTest{})

func (s *ConfigTest) SetUpSuite(c *C) {
}

func (s *ConfigTest) TearDownSuite(c *C) {
}

func (s *ConfigTest) TestConfig(c *C) {
	c.Check(GetPort(), Equals, uint16(1628))
	c.Check(GetAuth(), Equals, "auth-7609AF5AD2A3C6B2")
}
