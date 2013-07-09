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
func Literal(expr interface{}) *LiteralNode {
  literal := new(LiteralNode)
  literal.Expr = expr
  return literal
}

// OnNode factory method.
func On(expr interface{}) *OnNode {
  on := new(OnNode)
  on.Expr = expr
  return on
}

// LimitNode factory method.
func Limit(expr interface{}) *LimitNode {
  limit := new(LimitNode)
  limit.Expr = expr
  return limit
}

// OffsetNode factory method.
func Offset(expr interface{}) *OffsetNode {
  offset := new(OffsetNode)
  offset.Expr = expr
  return offset
}

// GroupNode factory method.
func Group(expr interface{}) *GroupNode {
  group := new(GroupNode)
  group.Expr = expr
  return group
}

// HavingNode factory method.
func Having(expr interface{}) *HavingNode {
  having := new(HavingNode)
  having.Expr = expr
  return having
}

func UnqualifiedColumn(expr interface{}) *UnqualifiedColumnNode {
  column := new(UnqualifiedColumnNode)
  column.Expr = expr
  return column
}

// ColumnNode factory method.
func Column(expr interface{}) *ColumnNode {
  column := new(ColumnNode)
  column.Expr = expr
  return column
}

// StarNode factory method.
func Star() *StarNode {
  return new(StarNode)
}
