package nodes

type SelectCoreNode struct {
  *Node
  Source        *JoinSourceNode
  Projections   []interface{}
  Wheres        []interface{}
  Groups        []interface{}
  Having        interface{}
  SetQuantifier interface{}
  Top           interface{}
}

func (core *SelectCoreNode) Project(a ...interface{}) *SelectCoreNode {
  core.Projections = append(core.Projections, a...)
  return core
}

func (core *SelectCoreNode) Where(a ...interface{}) *SelectCoreNode {
  core.Wheres = append(core.Wheres, a...)
  return core
}

func (core *SelectCoreNode) Join(a ...interface{}) *SelectCoreNode {
  core.Source.Join(a)
  return core
}

func SelectCore(relation *RelationNode) *SelectCoreNode {
  return &SelectCoreNode{
    &Node{nil, nil},
    JoinSource(relation),
    []interface{}{},
    []interface{}{},
    []interface{}{},
    nil, nil, nil}
}
