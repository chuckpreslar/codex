package nodes

type ReferenceNode struct {
  *Node
  aliases []string
}

func (reference *ReferenceNode) Name() string {
  return reference.Left().(string)
}

func (reference *ReferenceNode) Aliases() []string {
  return reference.aliases
}

func (reference *ReferenceNode) AddAliases(aliases ...string) ReferenceInterface {
  reference.aliases = append(reference.aliases, aliases...)
  return reference
}

func Reference(name string, aliases ...string) *ReferenceNode {
  return &ReferenceNode{&Node{name, nil}, aliases}
}
