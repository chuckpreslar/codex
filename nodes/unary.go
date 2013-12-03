// Package nodes provides nodes to use in codex AST's.
package nodes

// UnaryNode struct.
type UnaryNode struct {
	Expr interface{} // Single leaf for the Unary node.
}

type LiteralNode UnaryNode           // LiteralNode is a UnaryNode struct.
type BindingNode UnaryNode           // BindingNode is a UnaryNode struct.
type OnNode UnaryNode                // OnNode is a UnaryNode struct.
type LimitNode UnaryNode             // LimitNode is a UnaryNode struct.
type OffsetNode UnaryNode            // OffsetNode is a UnaryNode struct.
type GroupNode UnaryNode             // GroupNode is a UnaryNode struct.
type HavingNode UnaryNode            // HavingNode is a UnaryNode struct.
type UnqualifiedColumnNode UnaryNode // UnqualifiedColumnNode is a UnaryNode struct.
type ColumnNode UnaryNode            // ColumnNode is a UnaryNode struct.
type StarNode UnaryNode              // StarNode is a Unary node struct.
type EngineNode UnaryNode            // EngineNode is a Unary node struct.
type IndexNameNode UnaryNode         // IndexNameNode is a Unary node struct.

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

// BindingNode factory method.
func Binding() *BindingNode {
	return new(BindingNode)
}

// EngineNode factory method.
func Engine(expr interface{}) (engine *EngineNode) {
	engine = new(EngineNode)
	engine.Expr = expr
	return
}

// IndexNameNode factory method.
func IndexName(expr interface{}) (index *IndexNameNode) {
	index = new(IndexNameNode)
	index.Expr = expr
	return
}

func (u *UnqualifiedColumnNode) Eq(other interface{}) *EqualNode {
	return Equal(u, other)
}

func (u *UnqualifiedColumnNode) Neq(other interface{}) *NotEqualNode {
	return NotEqual(u, other)
}

func (u *UnqualifiedColumnNode) Gt(other interface{}) *GreaterThanNode {
	return GreaterThan(u, other)
}

func (u *UnqualifiedColumnNode) Gte(other interface{}) *GreaterThanOrEqualNode {
	return GreaterThanOrEqual(u, other)
}

func (u *UnqualifiedColumnNode) Lt(other interface{}) *LessThanNode {
	return LessThan(u, other)
}

func (u *UnqualifiedColumnNode) Lte(other interface{}) *LessThanOrEqualNode {
	return LessThanOrEqual(u, other)
}

func (u *UnqualifiedColumnNode) Like(other interface{}) *LikeNode {
	return Like(u, other)
}

func (u *UnqualifiedColumnNode) Unlike(other interface{}) *UnlikeNode {
	return Unlike(u, other)
}
