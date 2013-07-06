package nodes

type SelectCore struct {
  Relation    *Relation
  Source      *JoinSource
  Projections []interface{}
  Wheres      []interface{}
}
