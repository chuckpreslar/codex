package nodes

type InsertStatement struct {
  Columns  []interface{}
  Values   []interface{}
  Relation *Relation
}
