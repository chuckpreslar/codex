package nodes

type ReferenceNode struct {
  *Node
  aliases []string
}

func (reference *ReferenceNode) GetName() string {
  return reference.GetLeft().(string)
}

func (reference *ReferenceNode) GetAliases() []string {
  return reference.aliases
}

func (reference *ReferenceNode) AddAliases(aliases ...string) {
  reference.aliases = append(reference.aliases, aliases...)
}

func Reference(name string, aliases ...string) *ReferenceNode {
  reference := &ReferenceNode{&Node{name, nil}, aliases}
  return reference
}
