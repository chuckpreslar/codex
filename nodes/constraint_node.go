// Package nodes provides nodes to use in codex AST's.
package nodes

// ConstraintNode is a specific Binary node.
type ConstraintNode struct {
	Columns []interface{} // Column the ConstraintNode affects.
	Options []interface{} // Options to apply to the ConstraintNode.
}

type NotNullNode ConstraintNode    // NotNullNode is a ConstraintNode struct.
type UniqueNode ConstraintNode     // UniqueNode is a ConstraintNode struct.
type PrimaryKeyNode ConstraintNode // PrimaryKeyNode is a ConstraintNode struct.
type ForeignKeyNode ConstraintNode // ForeignKeyNode is a ConstraintNode struct.
type CheckNode ConstraintNode      // CheckNode is a ConstraintNode struct.
type DefaultNode ConstraintNode    // DefaultNode is a ConstraintNode struct.

// ConstraintNode factory method.
func Constraint(columns []interface{}, options ...interface{}) (constraint *ConstraintNode) {
	constraint = new(ConstraintNode)
	constraint.Columns = columns
	constraint.Options = options
	return
}

// NotNullNode factory method.
func NotNull(columns []interface{}, options ...interface{}) (nnull *NotNullNode) {
	nnull = new(NotNullNode)
	nnull.Columns = columns
	nnull.Options = options
	return
}

// UniqueNode factory method.
func Unique(columns []interface{}, options ...interface{}) (unique *UniqueNode) {
	unique = new(UniqueNode)
	unique.Columns = columns
	unique.Options = options
	return
}

// PrimaryKeyNode factory method.
func PrimaryKey(columns []interface{}, options ...interface{}) (pkey *PrimaryKeyNode) {
	pkey = new(PrimaryKeyNode)
	pkey.Columns = columns
	pkey.Options = options
	return
}

// ForeignKeyNode factory method.
func ForeignKey(columns []interface{}, options ...interface{}) (fkey *ForeignKeyNode) {
	fkey = new(ForeignKeyNode)
	fkey.Columns = columns
	fkey.Options = options
	return
}

// CheckNode factory method.
func Check(columns []interface{}, options ...interface{}) (check *CheckNode) {
	check = new(CheckNode)
	check.Columns = columns
	check.Options = options
	return
}

// DefaultNode factory method.
func Default(columns []interface{}, options ...interface{}) (def *DefaultNode) {
	def = new(DefaultNode)
	def.Columns = columns
	def.Options = options
	return
}
