package nodes

type SelectCore struct {
  source      *Reference
  projections []Attribute
  wheres      []ComparatorNode
}
