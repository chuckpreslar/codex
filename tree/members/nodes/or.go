package nodes

type Or struct {
  *Binary
}

func (o *Or) Or(other interface{}) *Or {
  return &Or{&Binary{o, other}}
}

func (o *Or) And(other interface{}) *And {
  return &And{&Binary{o, other}}
}
