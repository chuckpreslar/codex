package nodes

type NodeInterface interface {
  node()
}

type ComparisonInterface interface {
  NodeInterface
  Attribute() AttributeInterface
  Value() interface{}
  Or(ComparisonInterface) ComparisonInterface
}

type ComparableInterface interface {
  NodeInterface
  Eq(interface{}) ComparisonInterface
  Neq(interface{}) ComparisonInterface
  Gt(interface{}) ComparisonInterface
  Gte(interface{}) ComparisonInterface
  Lt(interface{}) ComparisonInterface
  Lte(interface{}) ComparisonInterface
  Matches(interface{}) ComparisonInterface
  As(interface{}) ComparableInterface
}

type RelationInterface interface {
  NodeInterface
  Name() string
  Aliases() []string
  AddAliases(...string) RelationInterface
}

type AttributeInterface interface {
  ComparableInterface
  Name() string
  Relation() RelationInterface
}

type SqlFunctionInterface interface {
  ComparableInterface
  FunctionName() string
}
