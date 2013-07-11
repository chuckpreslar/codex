package nodes

// AttributeNode is a specific Binary node
type AttributeNode struct {
  Name     interface{}   // Name of the Attribute.
  Relation *RelationNode // Relation the attribute belongs to.
}

// Returns and Equal node containing a reference to the
// attribute and other
func (self *AttributeNode) Eq(other interface{}) *EqualNode {
  return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// attribute and other
func (self *AttributeNode) Neq(other interface{}) *NotEqualNode {
  return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// attribute and other
func (self *AttributeNode) Gt(other interface{}) *GreaterThanNode {
  return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// attribute and other
func (self *AttributeNode) Gte(other interface{}) *GreaterThanOrEqualNode {
  return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// attribute and other
func (self *AttributeNode) Lt(other interface{}) *LessThanNode {
  return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// attribute and other
func (self *AttributeNode) Lte(other interface{}) *LessThanOrEqualNode {
  return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// attribute and other
func (self *AttributeNode) Like(other interface{}) *LikeNode {
  return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// attribute and other
func (self *AttributeNode) Unlike(other interface{}) *UnlikeNode {
  return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// attribute and other
func (self *AttributeNode) Or(other interface{}) *GroupingNode {
  return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// attribute and other
func (self *AttributeNode) And(other interface{}) *GroupingNode {
  return Grouping(And(self, other))
}

// Returns and Ascending node containing a reference to the
// attribute
func (self *AttributeNode) Asc() *AscendingNode {
  return Ascending(self)
}

// Returns and Descending node containing a reference to the
// attribute
func (self *AttributeNode) Desc() *DescendingNode {
  return Descending(self)
}

// AttributeNode factory method.
func Attribute(name interface{}, relation *RelationNode) (attribute *AttributeNode) {
  attribute = new(AttributeNode)
  attribute.Name = name
  attribute.Relation = relation
  return
}
