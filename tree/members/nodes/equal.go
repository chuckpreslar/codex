package nodes

type Equal struct {
  *Binary
}

func (e *Equal) Or(other interface{}) *Or {
  return &Or{&Binary{e, other}}
}

func (e *Equal) And(other interface{}) *And {
  return &And{&Binary{e, other}}
}
