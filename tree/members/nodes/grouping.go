package nodes

type Grouping Unary

func (grouping *Grouping) Or(other interface{}) *Grouping {
  return &Grouping{&Or{grouping, other}}
}

func (grouping *Grouping) And(other interface{}) *Grouping {
  return &Grouping{&And{grouping, other}}
}