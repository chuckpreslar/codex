package nodes

type NotEqual struct {
  *Binary
}

func (n *NotEqual) Or(other interface{}) *Or {
  return &Or{&Binary{n, other}}
}

func (n *NotEqual) And(other interface{}) *And {
  return &And{&Binary{n, other}}
}
