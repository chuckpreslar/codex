package nodes

type Like struct {
  *Binary
}

func (l *Like) Or(other interface{}) *Or {
  return &Or{&Binary{l, other}}
}

func (l *Like) And(other interface{}) *And {
  return &And{&Binary{l, other}}
}
