package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
	"github.com/stretchr/testify/assert"
)

func TestNilImplementsAcceptor(t *testing.T) {
	var node = &base.Nil{}
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *base.Nil to implement interfaces.Acceptor")
}
