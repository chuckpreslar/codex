package nodes

type Unlike struct {
  *Binary
}

func (u *Unlike) Or(other interface{}) *Or {
  return &Or{&Binary{u, other}}
}

func (u *Unlike) And(other interface{}) *And {
  return &And{&Binary{u, other}}
}
