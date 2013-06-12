package librarian

type Node struct {
  left  interface{}
  right interface{}
}

type EqualityNode interface {
  equality()
}

type EqNode struct {
  Node
}

type NeqNode struct {
  Node
}

type GtNode struct {
  Node
}

type GteNode struct {
  Node
}

type LtNode struct {
  Node
}

type LteNode struct {
  Node
}

type Matches struct {
  Node
}

func (n EqNode) equality()  {}
func (n NeqNode) equality() {}
func (n GtNode) equality()  {}
func (n GteNode) equality() {}
func (n LtNode) equality()  {}
func (n LteNode) equality() {}
func (n Matches) equality() {}

type JoinableNode interface {
  joinable()
}

type JoinNode struct {
  Node
}

type InnerJoinNode struct {
  Node
}

type OuterJoinNode struct {
  Node
}

type LeftJoinNode struct {
  Node
}

type RightJoinNode struct {
  Node
}

type FullJoinNode struct {
  Node
}

func (n JoinNode) joinable()      {}
func (n InnerJoinNode) joinable() {}
func (n OuterJoinNode) joinable() {}
func (n LeftJoinNode) joinable()  {}
func (n RightJoinNode) joinable() {}
func (n FullJoinNode) joinable()  {}
