package nodes

type SqlFunctionNode struct {
  *Node
}

func (function *SqlFunctionNode) FunctionName() string {
  return function.Right.(string)
}

func (function *SqlFunctionNode) Eq(value interface{}) ComparisonInterface {
  return Eq(function, value)
}

func (function *SqlFunctionNode) Neq(value interface{}) ComparisonInterface {
  return Neq(function, value)
}
func (function *SqlFunctionNode) Gt(value interface{}) ComparisonInterface {
  return Gt(function, value)
}

func (function *SqlFunctionNode) Gte(value interface{}) ComparisonInterface {
  return Gte(function, value)
}

func (function *SqlFunctionNode) Lt(value interface{}) ComparisonInterface {
  return Lt(function, value)
}

func (function *SqlFunctionNode) Lte(value interface{}) ComparisonInterface {
  return Lte(function, value)
}

func (function *SqlFunctionNode) Matches(value interface{}) ComparisonInterface {
  return Matches(function, value)
}

func Maximum(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "maximum"}}
}

func Minimum(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "minimum"}}
}

func Count(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "count"}}
}

func Average(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "average"}}
}

func First(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "first"}}
}

func Last(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "last"}}
}

func Sum(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "sum"}}
}

func UpperCase(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "upper"}}
}

func LowerCase(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "lower"}}
}

func Mid(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "mid"}}
}

func Length(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "length"}}
}

func Round(node NodeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{node, "round"}}
}
