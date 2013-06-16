package nodes

type NodeInterface interface {
  Left() interface{}
  Right() interface{}
}

type ComparisonInterface interface {
  Left() AttributeInterface
  Attribute() AttributeInterface
  Value() interface{}
  Or(ComparisonInterface) OrInterface
}

type ComparableInterface interface {
  Eq(interface{}) ComparisonInterface
  Neq(interface{}) ComparisonInterface
  Gt(interface{}) ComparisonInterface
  Gte(interface{}) ComparisonInterface
  Lt(interface{}) ComparisonInterface
  Lte(interface{}) ComparisonInterface
  Matches(interface{}) ComparisonInterface
}

type OrInterface interface {
  Left() ComparisonInterface
  Right() ComparisonInterface
}

type ReferenceInterface interface {
  NodeInterface
  GetName() string
  GetAliases() []string
}

type AttributeInterface interface {
  NodeInterface
  ComparableInterface
  GetName() string
  GetReference() ReferenceInterface
  GetTableName() string
}
