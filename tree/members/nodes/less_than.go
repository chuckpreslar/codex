package nodes

type LessThan struct {
  *Binary
}

func (l *LessThan) Or(other interface{}) *Or {
  return &Or{&Binary{l, other}}
}

func (l *LessThan) And(other interface{}) *And {
  return &And{&Binary{l, other}}
}
