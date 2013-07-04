package nodes

type Equal struct {
  *Binary
}

func (e *Equal) Or(other interface{}) *Or {
  return &Or{&Binary{e, other}}
}
