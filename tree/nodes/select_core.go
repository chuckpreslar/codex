package nodes

type SelectCore struct {
  Souce       *JoinSource
  Projections []interface{}
  Wheres      []interface{}
}
