package librarian

import (
  "fmt"
)

type column string

func (c column) Eq(value interface{}) *EqNode {
  eq := new(EqNode)
  eq.left = column
  eq.right = processEqualityValue(value)
  return eq
}

func (c column) Neq(value interface{}) *NeqNode {
  neq := new(NeqNode)
  neq.left = column
  neq.right = processEqualityValue(value)
  return neq
}

func (c column) Gt(value interface{}) *GtNode {
  gt := new(GtNode)
  gt.left = column
  gt.right = processEqualityValue(value)
  return gt
}

func (c column) Gte(value interface{}) *GteNode {
  gte := new(GteNode)
  gte.left = column
  gte.right = processEqualityValue(value)
  return gte
}

func (c column) Lt(value interface{}) *LtNode {
  lt := new(LtNode)
  lt.left = column
  lt.right = processEqualityValue(value)
  return lt
}

func (c column) Lte(value interface{}) *LteNode {
  lte := new(LteNode)
  lte.left = column
  lte.right = processEqualityValue(value)
  return lte
}

func (c column) Matches(value interface{}) *MatchesNode {
  matches := new(MatchesNode)
  matches.left = column
  matches.right = processEqualityValue(value)
  return matches
}

func processEqualityValue(value interface{}) interface{} {
  switch value.(type) {
  case int:
    return value
  case string:
    return fmt.Sprintf(`'%v'`, value)
  }
  return value
}
