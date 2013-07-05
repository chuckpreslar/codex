package nodes

type GreaterThan struct {
  *Binary
}

func (g *GreaterThan) Or(other interface{}) *Or {
  return &Or{&Binary{g, other}}
}

func (g *GreaterThan) And(other interface{}) *And {
  return &And{&Binary{g, other}}
}
