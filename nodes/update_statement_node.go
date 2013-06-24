package nodes

type UpdateStatementNode struct {
  *Node
  Relation *RelationNode
  Wheres   []interface{}
  Values   []interface{}
  Orders   []interface{}
  Limit    interface{}
  Key      interface{}
}

func UpdateStatement(relation *RelationNode) *UpdateStatementNode {
  return &UpdateStatementNode{&Node{}, relation, []interface{}{},
    []interface{}{}, []interface{}{}, nil, nil}
}
