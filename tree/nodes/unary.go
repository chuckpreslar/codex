package nodes

// UnaryNode struct
type UnaryNode struct {
  Expr interface{} // Single leaf for the Unary node.
}

type LiteralNode UnaryNode           // LiteralNode is a UnaryNode struct
type OnNode UnaryNode                // OnNode is a UnaryNode struct
type LimitNode UnaryNode             // LimitNode is a UnaryNode struct
type OffsetNode UnaryNode            // OffsetNode is a UnaryNode struct
type GroupNode UnaryNode             // GroupNode is a UnaryNode struct
type HavingNode UnaryNode            // HavingNode is a UnaryNode struct
type UnqualifiedColumnNode UnaryNode // UnqualifiedColumnNode is a UnaryNode struct
type ColumnNode UnaryNode            // ColumnNode is a UnaryNode struct
type StarNode UnaryNode              // Star node is a Unary node struct

// LiteralNode factory method.
func Literal(expr interface{}) (literal *LiteralNode) {
  literal = new(LiteralNode)
  literal.Expr = expr
  return
}

// OnNode factory method.
func On(expr interface{}) (on *OnNode) {
  on = new(OnNode)
  on.Expr = expr
  return
}

// LimitNode factory method.
func Limit(expr interface{}) (limit *LimitNode) {
  limit = new(LimitNode)
  limit.Expr = expr
  return
}

// OffsetNode factory method.
func Offset(expr interface{}) (offset *OffsetNode) {
  offset = new(OffsetNode)
  offset.Expr = expr
  return
}

// GroupNode factory method.
func Group(expr interface{}) (group *GroupNode) {
  group = new(GroupNode)
  group.Expr = expr
  return
}

// HavingNode factory method.
func Having(expr interface{}) (having *HavingNode) {
  having = new(HavingNode)
  having.Expr = expr
  return
}

func UnqualifiedColumn(expr interface{}) (column *UnqualifiedColumnNode) {
  column = new(UnqualifiedColumnNode)
  column.Expr = expr
  return
}

// ColumnNode factory method.
func Column(expr interface{}) (column *ColumnNode) {
  column = new(ColumnNode)
  column.Expr = expr
  return
}

// StarNode factory method.
func Star() *StarNode {
  return new(StarNode)
}
