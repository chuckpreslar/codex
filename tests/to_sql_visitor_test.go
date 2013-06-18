package tests

import (
  "fmt"
  "librarian/nodes"
  "librarian/visitors"
  "testing"
)

var visitor = visitors.ToSql()

func TestVisitEqualsNode(t *testing.T) {
  var right, left = 1, 2
  eq := nodes.Equals(right, left)
  sql := visitor.Accept(eq)
  exepcted := fmt.Sprintf("%v = %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitNotEqualsNode(t *testing.T) {
  var right, left = 1, 2
  neq := nodes.NotEquals(right, left)
  sql := visitor.Accept(neq)
  exepcted := fmt.Sprintf("%v != %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitGreaterThanNode(t *testing.T) {
  var right, left = 1, 2
  gt := nodes.GreaterThan(right, left)
  sql := visitor.Accept(gt)
  exepcted := fmt.Sprintf("%v > %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitGreaterThanOrEqualsNode(t *testing.T) {
  var right, left = 1, 2
  gte := nodes.GreaterThanOrEquals(right, left)
  sql := visitor.Accept(gte)
  exepcted := fmt.Sprintf("%v >= %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitLessThanNode(t *testing.T) {
  var right, left = 1, 2
  lt := nodes.LessThan(right, left)
  sql := visitor.Accept(lt)
  exepcted := fmt.Sprintf("%v < %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitLessThanOrEqualsNode(t *testing.T) {
  var right, left = 1, 2
  lte := nodes.LessThanOrEquals(right, left)
  sql := visitor.Accept(lte)
  exepcted := fmt.Sprintf("%v <= %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitAsNode(t *testing.T) {
  var right, left = 1, 2
  as := nodes.As(right, left)
  sql := visitor.Accept(as)
  exepcted := fmt.Sprintf("%v AS %v", right, left)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitOrNode(t *testing.T) {
  var right, left = 1, 2
  as := nodes.Equals(left, right).Or(nodes.Equals(left, right))
  sql := visitor.Accept(as)
  exepcted := fmt.Sprintf("%v = %v OR %v = %v", left, right, left, right)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}

func TestVisitAndNode(t *testing.T) {
  var right, left = 1, 2
  as := nodes.Equals(left, right).And(nodes.Equals(left, right))
  sql := visitor.Accept(as)
  exepcted := fmt.Sprintf("%v = %v AND %v = %v", left, right, left, right)
  if sql != exepcted {
    t.Errorf("EqNode: Expected %s, got %s\n.", exepcted, sql)
  }
}
