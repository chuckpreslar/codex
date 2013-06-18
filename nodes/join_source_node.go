package nodes

type JoinSourceNode struct {
  *Node
  Right []interface{}
}

func (source *JoinSourceNode) Join(a interface{}) *JoinSourceNode {
  source.Right = append(source.Right, a)
  return source
}

func JoinSource(relation *RelationNode) *JoinSourceNode {
  return &JoinSourceNode{&Node{relation, nil}, []interface{}{}}
}
