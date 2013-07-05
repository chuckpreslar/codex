package members

import (
  "librarian/tree/members"
  "librarian/tree/members/nodes"
  "testing"
)

func TestAttribute(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  if relation != attr.Relation {
    t.Errorf("Expected Attributes relation reference to be stored.")
  } else if name != attr.Name {
    t.Errorf("Expected Attributes name to be %v, got %v.", name, attr.Name)
  }
}

func TestAttributeEq(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  eq := attr.Eq(other)
  if attr != eq.Left {
    t.Errorf("Expect Left Equal leaf to equal %v, got %v.", attr, eq.Left)
  } else if other != eq.Right {
    t.Errorf("Expect Right Equal leaf to equal %v, got %v.", other, eq.Right)
  }
}

func TestAttributeNeq(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  neq := attr.Neq(other)
  if attr != neq.Left {
    t.Errorf("Expect Left NotEqual leaf to equal %v, got %v.", attr, neq.Left)
  } else if other != neq.Right {
    t.Errorf("Expect Right NotEqual leaf to equal %v, got %v.", other, neq.Right)
  }
}

func TestAttributeGt(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  gt := attr.Gt(other)
  if attr != gt.Left {
    t.Errorf("Expect Left GreaterThan leaf to equal %v, got %v.", attr, gt.Left)
  } else if other != gt.Right {
    t.Errorf("Expect Right GreaterThan leaf to equal %v, got %v.", other, gt.Right)
  }
}

func TestAttributeGte(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  gte := attr.Gte(other)
  if attr != gte.Left {
    t.Errorf("Expect Left GreaterThanOrEqual leaf to equal %v, got %v.", attr, gte.Left)
  } else if other != gte.Right {
    t.Errorf("Expect Right GreaterThanOrEqual leaf to equal %v, got %v.", other, gte.Right)
  }
}

func TestAttributeLt(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  lt := attr.Lt(other)
  if attr != lt.Left {
    t.Errorf("Expect Left LessThan leaf to equal %v, got %v.", attr, lt.Left)
  } else if other != lt.Right {
    t.Errorf("Expect Right LessThan leaf to equal %v, got %v.", other, lt.Right)
  }
}

func TestAttributeLte(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  lte := attr.Lte(other)
  if attr != lte.Left {
    t.Errorf("Expect Left LessThanOrEqual leaf to equal %v, got %v.", attr, lte.Left)
  } else if other != lte.Right {
    t.Errorf("Expect Right LessThanOrEqual leaf to equal %v, got %v.", other, lte.Right)
  }
}

func TestAttributeLike(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  like := attr.Like(other)
  if attr != like.Left {
    t.Errorf("Expect Left Like leaf to equal %v, got %v.", attr, like.Left)
  } else if other != like.Right {
    t.Errorf("Expect Right Like leaf to equal %v, got %v.", other, like.Right)
  }
}

func TestAttributeUnlike(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  unlike := attr.Unlike(other)
  if attr != unlike.Left {
    t.Errorf("Expect Left Like leaf to equal %v, got %v.", attr, unlike.Left)
  } else if other != unlike.Right {
    t.Errorf("Expect Right Like leaf to equal %v, got %v.", other, unlike.Right)
  }
}

func TestAttributeOr(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  or := attr.Or(other)
  if attr != or.Expr.(*nodes.Or).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", attr, or.Expr.(*nodes.Or).Left)
  } else if other != or.Expr.(*nodes.Or).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, or.Expr.(*nodes.Or).Right)
  }
}

func TestAttributeAnd(t *testing.T) {
  relation := &members.Relation{}
  name := "testing"
  attr := &members.Attribute{name, relation}
  other := 1
  and := attr.And(other)
  if attr != and.Expr.(*nodes.And).Left {
    t.Errorf("Expect Left Or leaf to equal %v, got %v.", attr, and.Expr.(*nodes.And).Left)
  } else if other != and.Expr.(*nodes.And).Right {
    t.Errorf("Expect Right Or leaf to equal %v, got %v.", other, and.Expr.(*nodes.And).Right)
  }
}
