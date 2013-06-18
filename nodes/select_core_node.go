package nodes

type SelectCoreNode struct {
  *Node
  Relation    *RelationNode
  Projections []interface{}
  Wheres      []interface{}
}

func (core *SelectCoreNode) Project(a ...interface{}) *SelectCoreNode {
  core.Projections = append(core.Projections, a...)
  return core
}

func (core *SelectCoreNode) Where(a ...interface{}) *SelectCoreNode {
  core.Wheres = append(core.Wheres, a...)
  return core
}

func SelectCore(relation *RelationNode) *SelectCoreNode {
  return &SelectCoreNode{&Node{nil, nil}, relation, []interface{}{}, []interface{}{}}
}
