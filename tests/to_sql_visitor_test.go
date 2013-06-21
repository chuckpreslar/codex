package tests

import (
  "fmt"
  "librarian"
  "librarian/nodes"
  "librarian/visitors"
  "strings"
  "testing"
)

var visitor = visitors.ToSql()

func TestVisitEqualsNode(t *testing.T) {
  var right, left = 1, 2
  eq := nodes.Equals(right, left)
  sql := strings.Trim(visitor.Accept(eq), " ")
  expected := fmt.Sprintf("%v = %v", right, left)
  if sql != expected {
    t.Errorf("EqualsNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitNotEqualsNode(t *testing.T) {
  var right, left = 1, 2
  neq := nodes.NotEquals(right, left)
  sql := strings.Trim(visitor.Accept(neq), " ")
  expected := fmt.Sprintf("%v != %v", right, left)
  if sql != expected {
    t.Errorf("NotEqualsNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitGreaterThanNode(t *testing.T) {
  var right, left = 1, 2
  gt := nodes.GreaterThan(right, left)
  sql := strings.Trim(visitor.Accept(gt), " ")
  expected := fmt.Sprintf("%v > %v", right, left)
  if sql != expected {
    t.Errorf("GreaterThanNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitGreaterThanOrEqualsNode(t *testing.T) {
  var right, left = 1, 2
  gte := nodes.GreaterThanOrEquals(right, left)
  sql := strings.Trim(visitor.Accept(gte), " ")
  expected := fmt.Sprintf("%v >= %v", right, left)
  if sql != expected {
    t.Errorf("GreaterThanOrEqualsNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitLessThanNode(t *testing.T) {
  var right, left = 1, 2
  lt := nodes.LessThan(right, left)
  sql := strings.Trim(visitor.Accept(lt), " ")
  expected := fmt.Sprintf("%v < %v", right, left)
  if sql != expected {
    t.Errorf("LessThanNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitLessThanOrEqualsNode(t *testing.T) {
  var right, left = 1, 2
  lte := nodes.LessThanOrEquals(right, left)
  sql := strings.Trim(visitor.Accept(lte), " ")
  expected := fmt.Sprintf("%v <= %v", right, left)
  if sql != expected {
    t.Errorf("LessThanOrEqualsNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitAsNode(t *testing.T) {
  var right, left = 1, 2
  as := nodes.As(right, left)
  sql := strings.Trim(visitor.Accept(as), " ")
  expected := fmt.Sprintf("%v AS %v", right, left)
  if sql != expected {
    t.Errorf("AsNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitOrNode(t *testing.T) {
  var right, left = 1, 2
  as := nodes.Equals(left, right).Or(nodes.Equals(left, right))
  sql := strings.Trim(visitor.Accept(as), " ")
  expected := fmt.Sprintf("%v = %v OR %v = %v", left, right, left, right)
  if sql != expected {
    t.Errorf("OrNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitAndNode(t *testing.T) {
  var right, left = 1, 2
  as := nodes.Equals(left, right).And(nodes.Equals(left, right))
  sql := strings.Trim(visitor.Accept(as), " ")
  expected := fmt.Sprintf("%v = %v AND %v = %v", left, right, left, right)
  if sql != expected {
    t.Errorf("AndNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitSelectCoreNode(t *testing.T) {
  table := librarian.NewTable("table")
  other := librarian.NewTable("other")
  core := nodes.SelectCore(table.Relation()).
    Project(table("col_one"), table("col_two")).
    Where(table("col_one").Equals("test")).
    Join(nodes.InnerJoin(other.On(other("id").Equals(table("other_id")))))
  sql := strings.Trim(visitor.Accept(core), " ")
  expected := fmt.Sprintf(`SELECT "table"."col_one", "table"."col_two" FROM "table" INNER JOIN "other" ON "other"."id" = "table"."other_id" WHERE "table"."col_one" = 'test'`)
  if sql != expected {
    t.Errorf("SelectCoreNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitOnNode(t *testing.T) {
  table := librarian.NewTable("table")
  other := librarian.NewTable("other")
  on := table.On(other("id").Equals(table("others_id")))
  sql := strings.Trim(visitor.Accept(on), " ")
  expected := fmt.Sprintf(`"table" ON "other"."id" = "table"."others_id"`)
  if sql != expected {
    t.Errorf("OnNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitInnerJoinNode(t *testing.T) {
  table := librarian.NewTable("table")
  other := librarian.NewTable("other")
  join := nodes.InnerJoin(table.On(other("id").Equals(table("others_id"))))
  sql := strings.Trim(visitor.Accept(join), " ")
  expected := fmt.Sprintf(`INNER JOIN "table" ON "other"."id" = "table"."others_id"`)
  if sql != expected {
    t.Errorf("InnerJoinNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitLimitNode(t *testing.T) {
  take := 1
  limit := nodes.Limit(take)
  sql := strings.Trim(visitor.Accept(limit), " ")
  expected := fmt.Sprintf(`LIMIT %v`, take)
  if sql != expected {
    t.Errorf("LimitNode: Expected %s, got %s\n.", expected, sql)
  }
}

func TestVisitOffsetNode(t *testing.T) {
  skip := 1
  offset := nodes.Offset(skip)
  sql := strings.Trim(visitor.Accept(offset), " ")
  expected := fmt.Sprintf(`OFFSET %v`, skip)
  if sql != expected {
    t.Errorf("OffsetNode: Expected %s, got %s\n.", expected, sql)
  }
}
