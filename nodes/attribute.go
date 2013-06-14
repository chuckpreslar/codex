package nodes

type Attribute struct {
  BaseNode
  name      string
  reference *Reference
}

func (attr Attribute) project() {}

func (attr Attribute) Eq(value interface{}) (eq EqNode) {
  eq = EqNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func (attr Attribute) Neq(value interface{}) (neq NeqNode) {
  neq = NeqNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func (attr Attribute) Gt(value interface{}) (gt GtNode) {
  gt = GtNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func (attr Attribute) Gte(value interface{}) (gte GteNode) {
  gte = GteNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func (attr Attribute) Lt(value interface{}) (lt LtNode) {
  lt = LtNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func (attr Attribute) Lte(value interface{}) (lte LteNode) {
  lte = LteNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func (attr Attribute) Matches(value interface{}) (matches MatchesNode) {
  matches = MatchesNode{ComparatorNode{BaseNode{attr, value}}}
  return
}

func NewAttribute(name string, reference *Reference) Attribute {
  return Attribute{BaseNode{}, name, reference}
}
