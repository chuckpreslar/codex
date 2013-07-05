package nodes

type LessThanOrEqual struct {
  *Binary
}

func (l *LessThanOrEqual) Or(other interface{}) *Or {
  return &Or{&Binary{l, other}}
}

func (l *LessThanOrEqual) And(other interface{}) *And {
  return &And{&Binary{l, other}}
}
