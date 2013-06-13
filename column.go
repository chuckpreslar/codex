package librarian

type Column struct {
  association Table
  name        string
}

func (column Column) Eq(value interface{}) EqNode {
  return EqNode{column, value}
}
func (column Column) Neq(value interface{}) NeqNode {
  return NeqNode{column, value}
}
func (column Column) Gt(value interface{}) GtNode {
  return GtNode{column, value}
}
func (column Column) Gte(value interface{}) GteNode {
  return GteNode{column, value}
}
func (column Column) Lt(value interface{}) LtNode {
  return LtNode{column, value}
}
func (column Column) Lte(value interface{}) LteNode {
  return LteNode{column, value}
}
func (column Column) Matches(value interface{}) MatchesNode {
  return MatchesNode{column, value}
}
