package server

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ServerTest struct{}

var _ = Suite(&ServerTest{})

func (s *ServerTest) TestAdd(c *C) {
	c.Assert(Add(2, 4), Equals, 6)
}
