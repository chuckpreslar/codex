package librarian

type JoinableNode interface {
  joinable()
}

type InnerJoinNode struct {
  left  table
  right EqualityNode
}

type OuterJoinNode struct {
  left  table
  right EqualityNode
}

type RightJoinNode struct {
  left  table
  right EqualityNode
}

type LeftJoinNode struct {
  left  table
  right EqualityNode
}

type FullJoinNode struct {
  left  table
  right EqualityNode
}

func (n *InnerJoinNode) joinable() {}
func (n *OuterJoinNode) joinable() {}
func (n *RightJoinNode) joinable() {}
func (n *LeftJoinNode) joinable()  {}
func (n *FullJoinNode) joinable()  {}
