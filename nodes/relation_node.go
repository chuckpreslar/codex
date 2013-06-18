package nodes

type RelationNode struct {
  *Node
  name       string
  aliases    []string
  tableAlias string
  engine     interface{}
  primaryKey interface{}
}

func Relation(name string) *RelationNode {
  return &RelationNode{&Node{nil, nil}, name, []string{}, "", nil, nil}
}
