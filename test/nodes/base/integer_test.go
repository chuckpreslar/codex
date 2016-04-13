package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
	"github.com/stretchr/testify/assert"
)

func TestIntegerImplementsAcceptor(t *testing.T) {
	var node = &base.Integer{Value: 1}
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *base.Integer to implement interfaces.Acceptor")
}
