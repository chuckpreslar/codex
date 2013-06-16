package nodes

type NodeInterface interface {
  GetLeft() interface{}
  GetRight() interface{}
}

type ReferenceInterface interface {
  NodeInterface
  GetName() string
  GetAliases() []string
}

type AttributeInterface interface {
  NodeInterface
  GetName() string
  GetReference() ReferenceInterface
  GetTableName() string
}
