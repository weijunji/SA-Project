package tools

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestTools(t *testing.T) { TestingT(t) }

type ToolsTest struct{}

var _ = Suite(&ToolsTest{})

func (s *ToolsTest) TestFileExist(c *C) {
	c.Check(FileExist("go.mod") || FileExist("tools_test.go"), Equals, true)
	c.Check(FileExist("fileNotExist"), Equals, false)
}
