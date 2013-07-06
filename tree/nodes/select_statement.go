package nodes

type SelectStatement struct {
  Cores []*SelectCore
  Limit *Limit
  Offset *Offset
}