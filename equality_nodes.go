package librarian

type EqualityNode interface {
  equality()
}

type EqNode struct {
  left  column
  right interface{}
}

type NeqNode struct {
  left  column
  right interface{}
}

type GtNode struct {
  left  column
  right interface{}
}

type GteNode struct {
  left  column
  right interface{}
}

type LtNode struct {
  left  column
  right interface{}
}

type LteNode struct {
  left  column
  right interface{}
}

type MatchesNode struct {
  left  column
  right interface{}
}

func (n *EqNode) equality()  {}
func (n *NeqNode) equality() {}
func (n *GtNode) equality()  {}
func (n *GteNode) equality() {}
func (n *LtNode) equality()  {}
func (n *LteNode) equality() {}
func (n *Matches) equality() {}
