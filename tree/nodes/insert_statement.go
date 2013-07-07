package nodes

type InsertStatement struct {
  Relation *Relation
  Columns  []interface{}
  Values   []interface{}
}
