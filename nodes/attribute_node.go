package nodes

type AttributeNode struct {
  *Node
  Name     string
  Relation *RelationNode
}

func (attribute *AttributeNode) Equals(b interface{}) *ComparisonNode {
  return Equals(attribute, b)
}

func (attribute *AttributeNode) Eq(b interface{}) *ComparisonNode {
  return Equals(attribute, b)
}

func (attribute *AttributeNode) NotEquals(b interface{}) *ComparisonNode {
  return NotEquals(attribute, b)
}

func (attribute *AttributeNode) Neq(b interface{}) *ComparisonNode {
  return NotEquals(attribute, b)
}

func (attribute *AttributeNode) GreaterThan(b interface{}) *ComparisonNode {
  return GreaterThan(attribute, b)
}

func (attribute *AttributeNode) Gt(b interface{}) *ComparisonNode {
  return GreaterThan(attribute, b)
}

func (attribute *AttributeNode) GreaterThanOrEquals(b interface{}) *ComparisonNode {
  return GreaterThanOrEquals(attribute, b)
}

func (attribute *AttributeNode) Gte(b interface{}) *ComparisonNode {
  return GreaterThanOrEquals(attribute, b)
}

func (attribute *AttributeNode) LessThan(b interface{}) *ComparisonNode {
  return LessThan(attribute, b)
}

func (attribute *AttributeNode) Lt(b interface{}) *ComparisonNode {
  return LessThan(attribute, b)
}

func (attribute *AttributeNode) LessThanOrEquals(b interface{}) *ComparisonNode {
  return LessThanOrEquals(attribute, b)
}

func (attribute *AttributeNode) Lte(b interface{}) *ComparisonNode {
  return LessThanOrEquals(attribute, b)
}

func (attribute *AttributeNode) Matches(b interface{}) *ComparisonNode {
  return Matches(attribute, b)
}

func (attribute *AttributeNode) Like(b interface{}) *ComparisonNode {
  return Matches(attribute, b)
}

func (attribute *AttributeNode) DoesNotMatch(b interface{}) *ComparisonNode {
  return DoesNotMatch(attribute, b)
}

func (attribute *AttributeNode) NotLike(b interface{}) *ComparisonNode {
  return DoesNotMatch(attribute, b)
}

func Attribute(name string, relation *RelationNode) *AttributeNode {
  return &AttributeNode{&Node{}, name, relation}
}
