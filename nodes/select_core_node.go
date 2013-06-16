package nodes

type SelectCoreNode struct {
  *Node
  reference   ReferenceInterface
  projections []AttributeInterface
  wheres      []ComparisonInterface
}

func (core *SelectCoreNode) Reference() ReferenceInterface {
  return core.reference
}

func (core *SelectCoreNode) AppendToProjections(attribute AttributeInterface) {
  core.projections = append(core.projections, attribute)
}

func (core *SelectCoreNode) Projections() []AttributeInterface {
  return core.projections
}

func (core *SelectCoreNode) AppendToWheres(comparison ComparisonInterface) {
  core.wheres = append(core.wheres, comparison)
}

func (core *SelectCoreNode) Wheres() []ComparisonInterface {
  return core.wheres
}

func SelectCore(reference ReferenceInterface, projections ...AttributeInterface) *SelectCoreNode {
  return &SelectCoreNode{&Node{nil, nil}, reference, projections, []ComparisonInterface{}}
}
