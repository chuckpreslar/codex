package nodes

type AttributeNode struct {
  *Node
  TableName string
}

func (attribute *AttributeNode) Relation() RelationInterface {
  return attribute.Left.(RelationInterface)
}

func (attribute *AttributeNode) Name() string {
  return attribute.Right.(string)
}

func (attribute *AttributeNode) Eq(other interface{}) ComparisonInterface {
  return Eq(attribute, other)
}

func (attribute *AttributeNode) Neq(other interface{}) ComparisonInterface {
  return Neq(attribute, other)
}
func (attribute *AttributeNode) Gt(other interface{}) ComparisonInterface {
  return Gt(attribute, other)
}

func (attribute *AttributeNode) Gte(other interface{}) ComparisonInterface {
  return Gte(attribute, other)
}

func (attribute *AttributeNode) Lt(other interface{}) ComparisonInterface {
  return Lt(attribute, other)
}

func (attribute *AttributeNode) Lte(other interface{}) ComparisonInterface {
  return Lte(attribute, other)
}

func (attribute *AttributeNode) Matches(other interface{}) ComparisonInterface {
  return Matches(attribute, other)
}

func (attribute *AttributeNode) As(other interface{}) ComparableInterface {
  return As(attribute, other)
}

func Attribute(relation RelationInterface, name string) AttributeInterface {
  return &AttributeNode{&Node{relation, name}, relation.Name()}
}
