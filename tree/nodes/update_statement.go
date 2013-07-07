package nodes

type UpdateStatement struct {
  Relation *Relation
  Values []interface{}
  Wheres []interface{}
  Limit *Limit
}