package nodes

type SelectStatementNode struct {
  *Node
  Relation RelationInterface
  Cores    []*SelectCoreNode
  Limit    int
  Offset   int
}

func (stmt *SelectStatementNode) Take(take int) *SelectStatementNode {
  stmt.Limit = take
  return stmt
}

func (stmt *SelectStatementNode) Skip(skip int) *SelectStatementNode {
  stmt.Offset = skip
  return stmt
}

func (stmt *SelectStatementNode) CoreAtIndex(index int) *SelectCoreNode {
  if index > len(stmt.Cores) {
    return nil
  }
  return stmt.Cores[index]
}

func SelectStatement(relation RelationInterface) *SelectStatementNode {
  core := SelectCore(relation)
  return &SelectStatementNode{&Node{}, relation, []*SelectCoreNode{core}, 0, 0}
}
