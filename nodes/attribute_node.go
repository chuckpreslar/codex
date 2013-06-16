package nodes

type AttributeNode struct {
  *Node
}

func (attribute *AttributeNode) GetName() string {
  return attribute.Left().(string)
}

func (attribute *AttributeNode) GetReference() ReferenceInterface {
  return attribute.Right().(ReferenceInterface)
}

func (attribute *AttributeNode) GetTableName() string {
  return attribute.GetReference().GetName()
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
  attribute := &AttributeNode{&Node{name, reference}}
  return attribute
}
