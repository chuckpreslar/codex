package nodes

type Reference struct {
  BaseNode
  name    string
  aliases []string
}

func NewReference(name string) (ref *Reference) {
  ref = new(Reference)
  ref.name = name
  return
}
