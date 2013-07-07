package nodes

type DeleteStatement struct {
  Relation *Relation
  Wheres   []interface{}
}