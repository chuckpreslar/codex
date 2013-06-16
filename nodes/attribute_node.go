package nodes

type AttributeNode struct {
  *Node
}

func (attribute *AttributeNode) Name() string {
  return attribute.Left().(string)
}

func (attribute *AttributeNode) Reference() ReferenceInterface {
  return attribute.Right().(ReferenceInterface)
}

func (attribute *AttributeNode) TableName() string {
  return attribute.Reference().Name()
}

func (attribute *AttributeNode) Eq(value interface{}) ComparisonInterface {
  return Eq(attribute, value)
}

func (attribute *AttributeNode) Neq(value interface{}) ComparisonInterface {
  return Neq(attribute, value)
}
func (attribute *AttributeNode) Gt(value interface{}) ComparisonInterface {
  return Gt(attribute, value)
}

func (attribute *AttributeNode) Gte(value interface{}) ComparisonInterface {
  return Gte(attribute, value)
}

func (attribute *AttributeNode) Lt(value interface{}) ComparisonInterface {
  return Lt(attribute, value)
}

func (attribute *AttributeNode) Lte(value interface{}) ComparisonInterface {
  return Lte(attribute, value)
}

func (attribute *AttributeNode) Matches(value interface{}) ComparisonInterface {
  return Matches(attribute, value)
}

func Attribute(name string, reference ReferenceInterface) *AttributeNode {
  return &AttributeNode{&Node{name, reference}}
}
