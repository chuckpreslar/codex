package nodes

type SelectCoreNode struct {
  *Node
  Relation    RelationInterface
  Projections []AttributeInterface
  Wheres      []ComparisonInterface
}

func (core *SelectCoreNode) AppendToProjections(attribute AttributeInterface) *SelectCoreNode {
  core.Projections = append(core.Projections, attribute)
  return core
}

func (core *SelectCoreNode) AppendToWheres(comparison ComparisonInterface) *SelectCoreNode {
  core.Wheres = append(core.Wheres, comparison)
  return core
}

func SelectCore(relation RelationInterface, projections ...AttributeInterface) *SelectCoreNode {
  return &SelectCoreNode{&Node{nil, nil}, relation, projections, []ComparisonInterface{}}
}
