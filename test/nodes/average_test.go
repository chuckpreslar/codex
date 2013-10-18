package nodes

import (
	"github.com/chuckpreslar/codex/nodes"
	"testing"
)

func TestCounter(t *testing.T) {
	avg := nodes.Average(1, 2, 3, 4)

	// The following struct members should exist.
	_ = avg.Expressions
	_ = avg.Alias
	_ = avg.Distinct

	// The following receiver methods should exist.
	_ = avg.And(1)
	_ = avg.Or(1)
	_ = avg.Eq(1)
	_ = avg.Neq(1)
	_ = avg.Gt(1)
	_ = avg.Gte(1)
	_ = avg.Lt(1)
	_ = avg.Lte(1)
	_ = avg.Like(1)
	_ = avg.Unlike(1)
}
