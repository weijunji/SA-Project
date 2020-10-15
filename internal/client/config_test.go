package client

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

func (s *ConfigTest) TestGenUUID(c *C) {
	c.Assert(generateUUID(), Not(Equals), generateUUID())
}

func (s *ConfigTest) TestConfig(c *C) {
	c.Check(GetHeartbeat(), Equals, 2)
	c.Check(GetAuth(), Equals, "auth-7609AF5AD2A3C6B2")
}
