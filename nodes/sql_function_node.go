package nodes

type SqlFunctionNode struct {
  *Node
}

func (function *SqlFunctionNode) FunctionName() string {
  return function.Right().(string)
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

func Maximum(attribute AttributeInterface) SqlFunctionInterface {
  return &SqlFunctionNode{&Node{attribute, "maximum"}}
}
