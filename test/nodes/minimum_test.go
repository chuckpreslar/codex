package nodes

import (
	"github.com/chuckpreslar/codex/nodes"
	"testing"
)

func TestMinimum(t *testing.T) {
	min := nodes.Minimum(1, 2, 3, 4)

	// The following struct members should exist.
	_ = min.Expressions
	_ = min.Alias
	_ = min.Distinct

	// The following receiver methods should exist.
	_ = min.And(1)
	_ = min.Or(1)
	_ = min.Eq(1)
	_ = min.Neq(1)
	_ = min.Gt(1)
	_ = min.Gte(1)
	_ = min.Lt(1)
	_ = min.Lte(1)
	_ = min.Like(1)
	_ = min.Unlike(1)
}
