package nodes

import "testing"

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
	"github.com/stretchr/testify/assert"
)

func TestAttributeImplementsAcceptor(t *testing.T) {
	var node = &nodes.Attribute{Name: "id", Relation: &nodes.Relation{Name: "users"}}
	assert.Implements(t, (*interfaces.Acceptor)(nil), node, "expected *node.Attribute to implement interfaces.Acceptor")
}
