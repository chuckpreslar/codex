package nodes

// AttributeNode is a specific Binary node
type AttributeNode struct {
  Name     interface{}   // Name of the Attribute.
  Relation *RelationNode // Relation the attribute belongs to.
}

// Returns and Equal node containing a reference to the
// attribute and other
func (attr *AttributeNode) Eq(other interface{}) *EqualNode {
  return Equal(attr, other)
}

// Returns and NotEqual node containing a reference to the
// attribute and other
func (attr *AttributeNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(attr, other)
}

// Returns and GreaterThan node containing a reference to the
// attribute and other
func (attr *AttributeNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(attr, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// attribute and other
func (attr *AttributeNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(attr, other)
}

// Returns and LessThan node containing a reference to the
// attribute and other
func (attr *AttributeNode) Lt(other interface{}) *LessThanNode {
  return LessThan(attr, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// attribute and other
func (attr *AttributeNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(attr, other)
}

// Returns and Like node containing a reference to the
// attribute and other
func (attr *AttributeNode) Like(other interface{}) *LikeNode {
  return Like(attr, other)
}

// Returns and Unlike node containing a reference to the
// attribute and other
func (attr *AttributeNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(attr, other)
}

// Returns and Or node containing a reference to the
// attribute and other
func (attr *AttributeNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(attr, other))
}

// Returns and And node containing a reference to the
// attribute and other
func (attr *AttributeNode) And(other interface{}) *GroupingNode {
  return Grouping(And(attr, other))
}

// Returns and Ascending node containing a reference to the
// attribute
func (attr *AttributeNode) Asc() *AscendingNode {
  return Ascending(attr)
}

// Returns and Descending node containing a reference to the
// attribute
func (attr *AttributeNode) Desc() *DescendingNode {
  return Descending(attr)
}

// AttributeNode factory method.
func Attribute(name interface{}, relation *RelationNode) *AttributeNode {
  attr := new(AttributeNode)
  attr.Name = name
  attr.Relation = relation
  return attr
}
