package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
	"github.com/stretchr/testify/assert"
)

func TestBooleanImplementsAcceptor(t *testing.T) {
	var node = &base.Boolean{}
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *base.Boolean to implement interfaces.Acceptor")
}
