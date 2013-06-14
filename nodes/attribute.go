package nodes

type Attribute struct {
  name      string
  reference *Reference
}

func (attr Attribute) Name() string {
  return attr.name
}

func (attr Attribute) SetName(name string) {
  attr.name = name
}

func (attr Attribute) Reference() *Reference {
  return attr.reference
}

func (attr Attribute) SetReference(reference *Reference) {
  attr.reference = reference
}

func NewAttribute(name string, reference *Reference) Attribute {
  return Attribute{name, reference}
}
