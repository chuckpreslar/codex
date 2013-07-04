package nodes

type And struct {
  *Binary
}

func (a *And) Or(other interface{}) *Or {
  return &Or{&Binary{a, other}}
}

func (a *And) And(other interface{}) *And {
  return &And{&Binary{a, other}}
}
