package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "github.com/chuckpreslar/codex/tree/visitors"
  "fmt"
  "testing"
)

func TestVisitEqual(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.Equal{left, right}
  expected := fmt.Sprintf("%v = %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitEqual was expected to return %s, got %s", expected, got)
  }

  equal = &nodes.Equal{left, nil}
  expected = fmt.Sprintf("%v IS NULL", left)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitEqual was expected to return %s, got %s", expected, got)
  }
}

func TestVisitNotEqual(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.NotEqual{left, right}
  expected := fmt.Sprintf("%v != %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitNotEqual was expected to return %s, got %s", expected, got)
  }

  equal = &nodes.NotEqual{left, nil}
  expected = fmt.Sprintf("%v IS NOT NULL", left)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitNotEqual was expected to return %s, got %s", expected, got)
  }
}

func TestVisitGreaterThan(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.GreaterThan{left, right}
  expected := fmt.Sprintf("%v > %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitGreaterThan was expected to return %s, got %s", expected, got)
  }
}

func TestVisitGreaterThanOrEqual(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.GreaterThanOrEqual{left, right}
  expected := fmt.Sprintf("%v >= %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitGreaterThanOrEqual was expected to return %s, got %s", expected, got)
  }
}

func TestVisitLessThan(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.LessThan{left, right}
  expected := fmt.Sprintf("%v < %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitLessThan was expected to return %s, got %s", expected, got)
  }
}

func TestVisitLessThanOrEqual(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.LessThanOrEqual{left, right}
  expected := fmt.Sprintf("%v <= %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitLessThanOrEqual was expected to return %s, got %s", expected, got)
  }
}

func TestVisitLike(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.Like{left, right}
  expected := fmt.Sprintf("%v LIKE %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitLike was expected to return %s, got %s", expected, got)
  }
}

func TestVisitUnlike(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.Unlike{left, right}
  expected := fmt.Sprintf("%v NOT LIKE %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitUnlike was expected to return %s, got %s", expected, got)
  }
}

func TestVisitOr(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.Or{left, right}
  expected := fmt.Sprintf("%v OR %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitOr was expected to return %s, got %s", expected, got)
  }
}

func TestVisitAnd(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.And{left, right}
  expected := fmt.Sprintf("%v AND %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitAnd was expected to return %s, got %s", expected, got)
  }
}

func TestVisitRelation(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  name, alias := "users", "site_members"
  equal := &nodes.Relation{name, ""}
  expected := fmt.Sprintf(`"%v"`, name)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitRelation was expected to return %s, got %s", expected, got)
  }
  equal.Alias = alias
  expected = fmt.Sprintf(`"%v"`, alias)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitRelation was expected to return %s, got %s", expected, got)
  }
}

func TestVisitAttribute(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  relationName, columnName := "testing", "id"
  relation := &nodes.Relation{relationName, ""}
  equal := &nodes.Attribute{columnName, relation}
  expected := fmt.Sprintf(`"%v"."%v"`, relationName, columnName)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitAttribute was expected to return %s, got %s", expected, got)
  }
}

func TestVisitGrouping(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  expr := 1
  equal := &nodes.Grouping{expr}
  expected := fmt.Sprintf("(%v)", expr)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitGrouping was expected to return %s, got %s", expected, got)
  }
}

func TestVisitNot(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  expr := 1
  equal := &nodes.Not{expr}
  expected := fmt.Sprintf("NOT (%v)", expr)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitNot was expected to return %s, got %s", expected, got)
  }
}

func TestVisitInnerJoin(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.InnerJoin{left, right}
  expected := fmt.Sprintf("INNER JOIN %v %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitInnerJoin was expected to return %s, got %s", expected, got)
  }
}

func TestVisitOuterJoin(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  left, right := 1, 1
  equal := &nodes.OuterJoin{left, right}
  expected := fmt.Sprintf("LEFT OUTER JOIN %v %v", left, right)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitInnerJoin was expected to return %s, got %s", expected, got)
  }
}

func TestVisitOn(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  expr := 1
  equal := &nodes.On{expr}
  expected := fmt.Sprintf("ON %v", expr)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitOn was expected to return %s, got %s", expected, got)
  }
}

func TestVisitLimit(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  expr := 1
  equal := &nodes.Limit{expr}
  expected := fmt.Sprintf(" LIMIT %v", expr)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitOn was expected to return %s, got %s", expected, got)
  }
}

func TestVisitOffset(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  expr := 1
  equal := &nodes.Offset{expr}
  expected := fmt.Sprintf(" OFFSET %v", expr)
  if got := visitor.Accept(equal); expected != got {
    t.Errorf("VisitOn was expected to return %s, got %s", expected, got)
  }
}

func TestVisitJoinSource(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  relation := &nodes.Relation{"testing", ""}
  inner := &nodes.InnerJoin{relation, &nodes.On{&nodes.Equal{1, 1}}}
  source := &nodes.JoinSource{relation, []interface{}{inner}}
  expected := `"testing" INNER JOIN "testing" ON 1 = 1`
  if got := visitor.Accept(source); expected != got {
    t.Errorf("VisitJoinSource was expected to return %s, got %s", expected, got)
  }
}

func TestVisitSelectCore(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  relation := &nodes.Relation{"testing", ""}
  attribute := &nodes.Attribute{"id", relation}
  inner := &nodes.InnerJoin{relation, &nodes.On{&nodes.Equal{1, 1}}}
  source := &nodes.JoinSource{relation, []interface{}{inner}}
  filter := &nodes.Grouping{&nodes.And{&nodes.Equal{1, 1}, &nodes.NotEqual{1, 2}}}
  core := &nodes.SelectCore{relation, source, []interface{}{attribute}, []interface{}{filter}}
  expected := fmt.Sprintf(`SELECT "testing"."id" FROM "testing" INNER JOIN "testing" ON 1 = 1 WHERE (1 = 1 AND 1 != 2)`)
  if got := visitor.Accept(core); expected != got {
    t.Errorf("VisitSelectCore was expected to return %s, got %s", expected, got)
  }
}

func TestVisitSelectStatement(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  relation := &nodes.Relation{"testing", ""}
  attribute := &nodes.Attribute{"id", relation}
  inner := &nodes.InnerJoin{relation, &nodes.On{&nodes.Equal{1, 1}}}
  source := &nodes.JoinSource{relation, []interface{}{inner}}
  filter := &nodes.Grouping{&nodes.And{&nodes.Equal{1, 1}, &nodes.NotEqual{1, 2}}}
  core := &nodes.SelectCore{relation, source, []interface{}{attribute}, []interface{}{filter}}
  statement := &nodes.SelectStatement{[]*nodes.SelectCore{core}, &nodes.Limit{1}, &nodes.Offset{1}}
  expected := fmt.Sprintf(`SELECT "testing"."id" FROM "testing" INNER JOIN "testing" ON 1 = 1 WHERE (1 = 1 AND 1 != 2) LIMIT 1 OFFSET 1`)
  if got := visitor.Accept(statement); expected != got {
    t.Errorf("VisitSelectStatement was expected to return %s, got %s", expected, got)
  }
}

func TestVisitInsertStatement(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  relation := &nodes.Relation{"testing", ""}
  statement := &nodes.InsertStatement{relation, []interface{}{"col", "umn"}, &nodes.Values{[]interface{}{"val", "ue"}, nil}}
  expected := `INSERT INTO "testing" ("col", "umn") VALUES ('val', 'ue')`
  if got := visitor.Accept(statement); expected != got {
    t.Errorf("VisitInsertStatement was expected to return %s, got %s", expected, got)
  }
}

func TestVisitString(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := `test`, `'test'`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitString was expected to return %s, got %s", expected, got)
  }
}

func TestVisitInteger(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := 0, `0`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitInteger was expected to return %s, got %s", expected, got)
  }
}

func TestVisitFloat(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := 0.25, `0.25`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitFloat was expected to return %s, got %s", expected, got)
  }
}

func TestVisitBool(t *testing.T) {
  visitor := &visitors.ToSqlVisitor{}
  value, expected := true, `'true'`
  if got := visitor.Accept(value); expected != got {
    t.Errorf("VisitBool was expected to return %s, got %s", expected, got)
  }
}
