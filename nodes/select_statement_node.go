package nodes

type SelectStatementNode struct {
  *Node
  Cores  []interface{}
  Limit  interface{}
  Lock   interface{}
  Offset interface{}
  Orders interface{}
  With   interface{}
}

func SelectStatement(relation *RelationNode) *SelectStatementNode {
  return &SelectStatementNode{&Node{},
    []interface{}{SelectCore(relation)},
    nil, nil, nil,
    nil, nil}
}
