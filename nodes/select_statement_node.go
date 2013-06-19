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

func SelectStatement() *SelectStatementNode {
  return &SelectStatementNode{&Node{}, []interface{}{}, nil, nil, nil, nil, nil}
}
