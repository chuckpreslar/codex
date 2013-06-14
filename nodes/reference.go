package nodes

type Reference struct {
  BaseNode
  name    string
  aliases []string
}

func (reference Reference) Name() string {
  return reference.name
}

func (reference Reference) Aliases() []string {
  return reference.aliases
}

func NewReference(name string) (ref *Reference) {
  ref = new(Reference)
  ref.name = name
  return
}
