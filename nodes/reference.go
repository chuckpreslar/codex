package nodes

type Reference struct {
  name    string
  aliases []string
}

func (reference *Reference) Name() string {
  return reference.name
}

func (reference *Reference) SetName(name string) {
  reference.name = name
}

func (reference *Reference) Aliases() []string {
  return reference.aliases
}

func (reference *Reference) AddAliases(aliases ...string) {
  reference.aliases = append(reference.aliases, aliases...)
}

func NewReference(name string, aliases ...string) *Reference {
  reference := new(Reference)
  reference.SetName(name)
  reference.AddAliases(aliases...)
  return reference
}
