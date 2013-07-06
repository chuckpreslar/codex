package nodes

// Unary node struct
type Unary struct {
  Expr interface{}
}

type Literal Unary
type On Unary
type Limit Unary
type Offset Unary
type Group Unary
type Having Unary
