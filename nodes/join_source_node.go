package nodes

type JoinSourceNode struct {
  *Node
  Right []interface{}
}

func JoinSource(relation *RelationNode) *JoinSourceNode {
  return &JoinSourceNode{&Node{relation, nil}, []interface{}{}}
}
