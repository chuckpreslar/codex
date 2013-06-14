package nodes

type Attribute struct {
  name      string
  reference *Reference
}

func (attr Attribute) Name() string {
  return attr.name
}

func (attr Attribute) SetName(name string) {
  attr.name = name
}

func (attr Attribute) Reference() *Reference {
  return attr.reference
}

func (attr Attribute) SetReference(reference *Reference) {
  attr.reference = reference
}

func (attr Attribute) Eq(value interface{}) ComparatorNode {
  return ComparatorNode{EqNode{attr, value}, nil}
}

func (attr Attribute) Neq(value interface{}) ComparatorNode {
  return ComparatorNode{NeqNode{attr, value}, nil}
}

func (attr Attribute) Gt(value interface{}) ComparatorNode {
  return ComparatorNode{GtNode{attr, value}, nil}
}

func (attr Attribute) Gte(value interface{}) ComparatorNode {
  return ComparatorNode{GteNode{attr, value}, nil}
}

func (attr Attribute) Lt(value interface{}) ComparatorNode {
  return ComparatorNode{LtNode{attr, value}, nil}
}

func (attr Attribute) Lte(value interface{}) ComparatorNode {
  return ComparatorNode{LteNode{attr, value}, nil}
}

func (attr Attribute) Matches(value interface{}) ComparatorNode {
  return ComparatorNode{MatchesNode{attr, value}, nil}
}

func NewAttribute(name string, reference *Reference) Attribute {
  return Attribute{name, reference}
}
