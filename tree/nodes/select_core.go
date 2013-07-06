package nodes

type SelectCore struct {
  Relation *Relation
  Projections []interface{}
  Wheres []interface{}
}