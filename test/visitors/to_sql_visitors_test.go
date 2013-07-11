package visitors

import (
  "github.com/chuckpreslar/codex/tree/nodes"
  "github.com/chuckpreslar/codex/tree/visitors"
  "testing"
)

var sql = &visitors.ToSqlVisitor{}

func TestAccept(t *testing.T) {
  got, err := sql.Accept(nil)
  if err == nil {
    t.Errorf("TestAccept was expected to return an error but did not, received %v, %v", got, err)
  }
}

func TestGrouping(t *testing.T) {
  grouping := nodes.Grouping(1)
  expected := "(1)"
  if got, _ := sql.Accept(grouping); expected != got {
    t.Errorf("TestGrouping was expected to return %s, got %s", expected, got)
  }
}

func TestAnd(t *testing.T) {
  and := nodes.And(1, 2)
  expected := "1 AND 2"
  if got, _ := sql.Accept(and); expected != got {
    t.Errorf("TestAnd was expected to return %s, got %s", expected, got)
  }
}

func TestOr(t *testing.T) {
  or := nodes.Or(1, 2)
  expected := "1 OR 2"
  if got, _ := sql.Accept(or); expected != got {
    t.Errorf("TestOr was expected to return %s, got %s", expected, got)
  }
}

func TestExtensiveGrouping(t *testing.T) {
  ext := nodes.And(1, 2).Or(3)
  expected := "(1 AND 2 OR 3)"
  if got, _ := sql.Accept(ext); expected != got {
    t.Errorf("TestExtensiveGrouping was expected to return %s, got %s", expected, got)
  }
}

func TestEqual(t *testing.T) {
  eq := nodes.Equal(1, 2)
  expected := "1 = 2"
  if got, _ := sql.Accept(eq); expected != got {
    t.Errorf("TestEqual was expected to return %s, got %s", expected, got)
  }
}

func TestNotEqual(t *testing.T) {
  neq := nodes.NotEqual(1, 2)
  expected := "1 != 2"
  if got, _ := sql.Accept(neq); expected != got {
    t.Errorf("TestNotEqual was expected to return %s, got %s", expected, got)
  }
}

func TestGreaterThan(t *testing.T) {
  gt := nodes.GreaterThan(1, 2)
  expected := "1 > 2"
  if got, _ := sql.Accept(gt); expected != got {
    t.Errorf("TestGreaterThan was expected to return %s, got %s", expected, got)
  }
}

func TestGreaterThanOrEqual(t *testing.T) {
  gte := nodes.GreaterThanOrEqual(1, 2)
  expected := "1 >= 2"
  if got, _ := sql.Accept(gte); expected != got {
    t.Errorf("TestGreaterThanOrEqual was expected to return %s, got %s", expected, got)
  }
}

func TestLessThan(t *testing.T) {
  lt := nodes.LessThan(1, 2)
  expected := "1 < 2"
  if got, _ := sql.Accept(lt); expected != got {
    t.Errorf("TestLessThan was expected to return %s, got %s", expected, got)
  }
}

func TestLessThanOrEqual(t *testing.T) {
  lte := nodes.LessThanOrEqual(1, 2)
  expected := "1 <= 2"
  if got, _ := sql.Accept(lte); expected != got {
    t.Errorf("TestLessThanOrEqual was expected to return %s, got %s", expected, got)
  }
}

func TestLike(t *testing.T) {
  like := nodes.Like(1, 2)
  expected := "1 LIKE 2"
  if got, _ := sql.Accept(like); expected != got {
    t.Errorf("TestLike was expected to return %s, got %s", expected, got)
  }
}

func TestUnlike(t *testing.T) {
  unlike := nodes.Unlike(1, 2)
  expected := "1 NOT LIKE 2"
  if got, _ := sql.Accept(unlike); expected != got {
    t.Errorf("TestUnlike was expected to return %s, got %s", expected, got)
  }
}

func TestExtensiveComparison(t *testing.T) {
  comparison := nodes.Equal(1, 2).Or(nodes.NotEqual(3, 4).And(5))
  expected := "(1 = 2 OR (3 != 4 AND 5))"
  if got, _ := sql.Accept(comparison); expected != got {
    t.Errorf("TestExtensiveComparison was expected to return %s, got %s", expected, got)
  }
}

func TestUnaliasedRelation(t *testing.T) {
  relation := nodes.Relation("table")
  expected := `"table"`
  if got, _ := sql.Accept(relation); expected != got {
    t.Errorf("TestUnaliasedRelation was expected to return %s, got %s", expected, got)
  }
}

func TestColumn(t *testing.T) {
  column := nodes.Column("column")
  expected := `"column"`
  if got, _ := sql.Accept(column); expected != got {
    t.Errorf("TestColumn was expected to return %s, got %s", expected, got)
  }
}

func TestUnaliasedAttribute(t *testing.T) {
  relation := nodes.Relation("table")
  column := nodes.Column("column")
  attribute := nodes.Attribute(column, relation)
  expected := `"table"."column"`
  if got, _ := sql.Accept(attribute); expected != got {
    t.Errorf("TestUnaliasedAttribute was expected to return %s, got %s", expected, got)
  }
}

func TestJoinSource(t *testing.T) {
  relation := nodes.Relation("table")
  source := nodes.JoinSource(relation)
  expected := `"table"`
  if got, _ := sql.Accept(source); expected != got {
    t.Errorf("TestJoinSource was expected to return %s, got %s", expected, got)
  }
}

func TestExtensiveJoinSource(t *testing.T) {
  relation := nodes.Relation("table")
  source := nodes.JoinSource(relation)
  source.Right = append(source.Right, []interface{}{1, 2, 3}...)
  expected := `"table" 1 2 3`
  if got, _ := sql.Accept(source); expected != got {
    t.Errorf("TestExtensiveJoinSource was expected to return %s, got %s", expected, got)
  }
}

func TestCount(t *testing.T) {
  count := nodes.Count(1)
  expected := "COUNT(1)"
  if got, _ := sql.Accept(count); expected != got {
    t.Errorf("TestCount was expected to return %s, got %s", expected, got)
  }
}

func TestSum(t *testing.T) {
  sum := nodes.Sum(1)
  expected := "SUM(1)"
  if got, _ := sql.Accept(sum); expected != got {
    t.Errorf("TestSum was expected to return %s, got %s", expected, got)
  }
}

func TestAverage(t *testing.T) {
  avg := nodes.Average(1)
  expected := "AVG(1)"
  if got, _ := sql.Accept(avg); expected != got {
    t.Errorf("TestAverage was expected to return %s, got %s", expected, got)
  }
}

func TestMinimum(t *testing.T) {
  min := nodes.Minimum(1)
  expected := "MIN(1)"
  if got, _ := sql.Accept(min); expected != got {
    t.Errorf("TestMinimum was expected to return %s, got %s", expected, got)
  }
}

func TestMaximum(t *testing.T) {
  max := nodes.Maximum(1)
  expected := "MAX(1)"
  if got, _ := sql.Accept(max); expected != got {
    t.Errorf("TestMaximum was expected to return %s, got %s", expected, got)
  }
}

func TestExtensiveFunction(t *testing.T) {
  function := nodes.Sum(1).Or(nodes.Count(2).And(nodes.Average(3)))
  expected := "(SUM(1) OR (COUNT(2) AND AVG(3)))"
  if got, _ := sql.Accept(function); expected != got {
    t.Errorf("TestExtensiveFunction was expected to return %s, got %s", expected, got)
  }
}

func TestLimit(t *testing.T) {
  limit := nodes.Limit(1)
  expected := "LIMIT 1"
  if got, _ := sql.Accept(limit); expected != got {
    t.Errorf("TestLimit was expected to return %s, got %s", expected, got)
  }
}

func TestOffset(t *testing.T) {
  offset := nodes.Offset(1)
  expected := "OFFSET 1"
  if got, _ := sql.Accept(offset); expected != got {
    t.Errorf("TestOffset was expected to return %s, got %s", expected, got)
  }
}

func TestHaving(t *testing.T) {
  having := nodes.Having(1)
  expected := "HAVING 1"
  if got, _ := sql.Accept(having); expected != got {
    t.Errorf("TestHaving was expected to return %s, got %s", expected, got)
  }
}

func TestOn(t *testing.T) {
  on := nodes.On(1)
  expected := "ON 1"
  if got, _ := sql.Accept(on); expected != got {
    t.Errorf("TestOn was expected to return %s, got %s", expected, got)
  }
}

func TestAscending(t *testing.T) {
  asc := nodes.Ascending(1)
  expected := "1 ASC"
  if got, _ := sql.Accept(asc); expected != got {
    t.Errorf("TestAscending was expected to return %s, got %s", expected, got)
  }
}

func TestDescending(t *testing.T) {
  asc := nodes.Descending(1)
  expected := "1 DESC"
  if got, _ := sql.Accept(asc); expected != got {
    t.Errorf("TestDescending was expected to return %s, got %s", expected, got)
  }
}

func TestInnerJoin(t *testing.T) {
  join := nodes.InnerJoin(1, nil)
  expected := "INNER JOIN 1"
  if got, _ := sql.Accept(join); expected != got {
    t.Errorf("TestInnerJoin was expected to return %s, got %s", expected, got)
  }
}

func TestOuterJoin(t *testing.T) {
  join := nodes.OuterJoin(1, 2)
  expected := "LEFT OUTER JOIN 1 2"
  if got, _ := sql.Accept(join); expected != got {
    t.Errorf("TestOuterJoin was expected to return %s, got %s", expected, got)
  }
}

func TestSelectCore(t *testing.T) {
  relation := nodes.Relation("table")
  core := nodes.SelectCore(relation)
  expected := `SELECT FROM "table"`
  if got, _ := sql.Accept(core); expected != got {
    t.Errorf("TestSelectCore was expected to return %s, got %s", expected, got)
  }
}

func TestSelectCoreExtensive(t *testing.T) {
  relation := nodes.Relation("table")
  core := nodes.SelectCore(relation)
  core.Projections = append(core.Projections, 1, 2)
  core.Wheres = append(core.Wheres, 3, 4)
  core.Source.Right = append(core.Source.Right, nodes.InnerJoin(5, nil))
  expected := `SELECT 1, 2 FROM "table" INNER JOIN 5 WHERE 3 AND 4`
  if got, _ := sql.Accept(core); expected != got {
    t.Errorf("TestSelectCoreExtensive was expected to return %s, got %s", expected, got)
  }
}

func TestSelectStatement(t *testing.T) {
  relation := nodes.Relation("table")
  stmt := nodes.SelectStatement(relation)
  expected := `SELECT FROM "table"`
  if got, _ := sql.Accept(stmt); expected != got {
    t.Errorf("TestSelectStatement was expected to return %s, got %s", expected, got)
  }
}

func TestInsertStatement(t *testing.T) {
  relation := nodes.Relation("table")
  stmt := nodes.InsertStatement(relation)
  expected := `INSERT INTO "table" `
  if got, _ := sql.Accept(stmt); expected != got {
    t.Errorf("TestInsertStatement was expected to return %s, got %s", expected, got)
  }
}

func TestUpdateStatement(t *testing.T) {
  relation := nodes.Relation("table")
  stmt := nodes.UpdateStatement(relation)
  expected := `UPDATE "table" `
  if got, _ := sql.Accept(stmt); expected != got {
    t.Errorf("TestUpdateStatement was expected to return %s, got %s", expected, got)
  }
}

func TestDeleteStatement(t *testing.T) {
  relation := nodes.Relation("table")
  stmt := nodes.DeleteStatement(relation)
  expected := `DELETE FROM "table" `
  if got, _ := sql.Accept(stmt); expected != got {
    t.Errorf("TestDeleteStatement was expected to return %s, got %s", expected, got)
  }
}

func TestAssignment(t *testing.T) {
  assignment := nodes.Assignment(1, 2)
  expected := "1 = 2"
  if got, _ := sql.Accept(assignment); expected != got {
    t.Errorf("TestAssignment was expected to return %s, got %s", expected, got)
  }
}

func TestLiteral(t *testing.T) {
  number := nodes.Literal(1)
  str := nodes.Literal("1")
  n, _ := sql.Accept(number)
  s, _ := sql.Accept(str)
  if n != s {
    t.Errorf("TestLiteral was expected to return the same result, got %s and %s", n, s)
  }
}

func TestUnqualifiedColumn(t *testing.T) {
  column := nodes.UnqualifiedColumn("column")
  expected := `"column"`
  if got, _ := sql.Accept(column); expected != got {
    t.Errorf("TestUnqualifiedColumn was expected to return %s, got %s", expected, got)
  }
}

func TestVisitString(t *testing.T) {
  value, expected := `test`, `'test'`
  if got, _ := sql.Accept(value); expected != got {
    t.Errorf("TestVisitString was expected to return %s, got %s", expected, got)
  }
}

func TestVisitInteger(t *testing.T) {
  value, expected := 0, `0`
  if got, _ := sql.Accept(value); expected != got {
    t.Errorf("TestVisitInteger was expected to return %s, got %s", expected, got)
  }
}

func TestVisitFloat(t *testing.T) {
  value, expected := 0.25, `0.25`
  if got, _ := sql.Accept(value); expected != got {
    t.Errorf("TestVisitFloat was expected to return %s, got %s", expected, got)
  }
}

func TestVisitBool(t *testing.T) {
  value, expected := true, `'true'`
  if got, _ := sql.Accept(value); expected != got {
    t.Errorf("TestVisitBool was expected to return %s, got %s", expected, got)
  }
}
