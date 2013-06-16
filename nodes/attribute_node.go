package nodes

type AttributeNode struct {
  *Node
}

func (attr *AttributeNode) GetName() string {
  return attr.GetLeft().(string)
}

func (attr *AttributeNode) GetReference() ReferenceInterface {
  return attr.GetRight().(ReferenceInterface)
}

func (attr *AttributeNode) GetTableName() string {
  return attr.GetReference().GetName()
}

func Attribute(name string, reference ReferenceInterface) *AttributeNode {
  attribute := &AttributeNode{&Node{name, reference}}
  return attribute
}
