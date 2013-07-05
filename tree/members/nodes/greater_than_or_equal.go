package nodes

type GreaterThanOrEqual struct {
  *Binary
}

func (g *GreaterThanOrEqual) Or(other interface{}) *Or {
  return &Or{&Binary{g, other}}
}

func (g *GreaterThanOrEqual) And(other interface{}) *And {
  return &And{&Binary{g, other}}
}
