package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
	"github.com/stretchr/testify/assert"
)

func TestRelationImplementsAcceptor(t *testing.T) {
	var node = &nodes.Relation{Name: "users"}
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *nodes.Relation to implement interfaces.Acceptor")
}
