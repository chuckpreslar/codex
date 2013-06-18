package nodes

type SqlFunctionNode struct {
  *Node
}

func (function *SqlFunctionNode) FunctionName() string {
  return function.Right.(string)
}

func (function *SqlFunctionNode) Eq(other interface{}) ComparisonInterface {
  return Eq(function, other)
}

func (function *SqlFunctionNode) Neq(other interface{}) ComparisonInterface {
  return Neq(function, other)
}
func (function *SqlFunctionNode) Gt(other interface{}) ComparisonInterface {
  return Gt(function, other)
}

func (function *SqlFunctionNode) Gte(other interface{}) ComparisonInterface {
  return Gte(function, other)
}

func (function *SqlFunctionNode) Lt(other interface{}) ComparisonInterface {
  return Lt(function, other)
}

func (function *SqlFunctionNode) Lte(other interface{}) ComparisonInterface {
  return Lte(function, other)
}

func (function *SqlFunctionNode) Matches(other interface{}) ComparisonInterface {
  return Matches(function, other)
}

func (function *SqlFunctionNode) As(other interface{}) ComparableInterface {
  return As(function, other)
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
