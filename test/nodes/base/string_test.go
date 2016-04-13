package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
	"github.com/stretchr/testify/assert"
)

func TestStringImplementsAcceptor(t *testing.T) {
	var node = &base.String{Value: "bob"}
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *base.String to implement interfaces.Acceptor")
}
