package nodes

import (
	"github.com/chuckpreslar/codex/nodes"
	"testing"
)

func TestAnd(t *testing.T) {
	and := nodes.And(1, 2)

	// The following struct members should exist.
	_ = and.Left
	_ = and.Right

	// The following receiver methods should exist.
	_ = and.Or(1)
	_ = and.And(1)
	_ = and.Not()
}
