package nodes

type Attribute struct {
  BaseNode
  name      string
  reference *Reference
}

func (attr Attribute) Eq(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{EqNode{BaseNode{attr, value}}, nil}}
}

func (attr Attribute) Neq(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{NeqNode{BaseNode{attr, value}}, nil}}
}

func (attr Attribute) Gt(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{GtNode{BaseNode{attr, value}}, nil}}
}

func (attr Attribute) Gte(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{GteNode{BaseNode{attr, value}}, nil}}
}

func (attr Attribute) Lt(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{LtNode{BaseNode{attr, value}}, nil}}
}

func (attr Attribute) Lte(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{LteNode{BaseNode{attr, value}}, nil}}
}

func (attr Attribute) Matches(value interface{}) ComparatorNode {
  return ComparatorNode{BaseNode{MatchesNode{BaseNode{attr, value}}, nil}}
}

func NewAttribute(name string, reference *Reference) Attribute {
  return Attribute{BaseNode{}, name, reference}
}
