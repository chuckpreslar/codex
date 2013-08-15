// Package nodes provides nodes to use in codex AST's.
package nodes

// ConstraintNode is a Nary node.
type ConstraintNode struct {
  Column  interface{}
  Options []interface{}
}

type NotNullNode ConstraintNode
type UniqueNode ConstraintNode
type PrimaryKeyNode ConstraintNode
type ForeignKeyNode ConstraintNode
type CheckNode ConstraintNode
type DefaultNode ConstraintNode

// ConstraintNode factory method.
func Constraint(column interface{}, options ...interface{}) (constraint *ConstraintNode) {
  constraint = new(ConstraintNode)
  constraint.Column = column
  constraint.Options = options
  return
}

func NotNull(column interface{}, options ...interface{}) (nnull *NotNullNode) {
  nnull = new(NotNullNode)
  nnull.Column = column
  nnull.Options = options
  return
}

func Unique(column interface{}, options ...interface{}) (unique *UniqueNode) {
  unique = new(UniqueNode)
  unique.Column = column
  unique.Options = options
  return
}

func PrimaryKey(column interface{}, options ...interface{}) (pkey *PrimaryKeyNode) {
  pkey = new(PrimaryKeyNode)
  pkey.Column = column
  pkey.Options = options
  return
}

func ForeignKey(column interface{}, options ...interface{}) (fkey *ForeignKeyNode) {
  fkey = new(ForeignKeyNode)
  fkey.Column = column
  fkey.Options = options
  return
}

func Check(column interface{}, options ...interface{}) (check *CheckNode) {
  check = new(CheckNode)
  check.Column = column
  check.Options = options
  return
}

func Default(column interface{}, options ...interface{}) (def *DefaultNode) {
  def = new(DefaultNode)
  def.Column = column
  def.Options = options
  return
}
