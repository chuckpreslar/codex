package lib

type Table func(string) *Attribute

func NewTable(name string) Table {
  relation := &Relation{name, ""}
  return func(name string) *Attribute {
    return &Attribute{name, relation}
  }
}
