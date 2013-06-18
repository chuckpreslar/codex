package nodes

type AttributeNode struct {
  *Node
  name     string
  relation *RelationNode
}

func Attribute(name string, relation *RelationNode) *AttributeNode {
  return &AttributeNode{&Node{}, name, relation}
}
