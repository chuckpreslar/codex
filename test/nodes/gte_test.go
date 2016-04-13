package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
	"github.com/stretchr/testify/assert"
)

func TestGteImplementsAcceptor(t *testing.T) {
	var (
		child = aux.CreateAcceptorFrom(1)
		node  = &nodes.Gte{Left: child, Right: child}
	)
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *node.Gte to implement interfaces.Acceptor")
}
