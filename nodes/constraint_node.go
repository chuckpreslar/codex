// Package nodes provides nodes to use in codex AST's.
package nodes

// ConstraintNode is a specific Binary node.
type ConstraintNode struct {
  Column  interface{}   // Column the ConstraintNode affects.
  Options []interface{} // Options to apply to the ConstraintNode.
}

type NotNullNode ConstraintNode    // NotNullNode is a ConstraintNode struct.
type UniqueNode ConstraintNode     // UniqueNode is a ConstraintNode struct.
type PrimaryKeyNode ConstraintNode // PrimaryKeyNode is a ConstraintNode struct.
type ForeignKeyNode ConstraintNode // ForeignKeyNode is a ConstraintNode struct.
type CheckNode ConstraintNode      // CheckNode is a ConstraintNode struct.
type DefaultNode ConstraintNode    // DefaultNode is a ConstraintNode struct.

// ConstraintNode factory method.
func Constraint(column interface{}, options ...interface{}) (constraint *ConstraintNode) {
  constraint = new(ConstraintNode)
  constraint.Column = column
  constraint.Options = options
  return
}

// NotNullNode factory method.
func NotNull(column interface{}, options ...interface{}) (nnull *NotNullNode) {
  nnull = new(NotNullNode)
  nnull.Column = column
  nnull.Options = options
  return
}

// UniqueNode factory method.
func Unique(column interface{}, options ...interface{}) (unique *UniqueNode) {
  unique = new(UniqueNode)
  unique.Column = column
  unique.Options = options
  return
}

// PrimaryKeyNode factory method.
func PrimaryKey(column interface{}, options ...interface{}) (pkey *PrimaryKeyNode) {
  pkey = new(PrimaryKeyNode)
  pkey.Column = column
  pkey.Options = options
  return
}

// ForeignKeyNode factory method.
func ForeignKey(column interface{}, options ...interface{}) (fkey *ForeignKeyNode) {
  fkey = new(ForeignKeyNode)
  fkey.Column = column
  fkey.Options = options
  return
}

// CheckNode factory method.
func Check(column interface{}, options ...interface{}) (check *CheckNode) {
  check = new(CheckNode)
  check.Column = column
  check.Options = options
  return
}

// DefaultNode factory method.
func Default(column interface{}, options ...interface{}) (def *DefaultNode) {
  def = new(DefaultNode)
  def.Column = column
  def.Options = options
  return
}
