package nodes

type SelectCoreNode struct {
  *Node
  Relation    RelationInterface
  Projections []ComparableInterface
  Wheres      []ComparisonInterface
}

func (core *SelectCoreNode) core() {}

func (core *SelectCoreNode) AppendToProjections(comparable ComparableInterface) *SelectCoreNode {
  core.Projections = append(core.Projections, comparable)
  return core
}

func (core *SelectCoreNode) AppendToWheres(comparison ComparisonInterface) *SelectCoreNode {
  core.Wheres = append(core.Wheres, comparison)
  return core
}

func SelectCore(relation RelationInterface, projections ...ComparableInterface) *SelectCoreNode {
  return &SelectCoreNode{&Node{nil, nil}, relation, projections, []ComparisonInterface{}}
}
