package managers

import (
  "github.com/chuckpreslar/codex/factory"
  "github.com/chuckpreslar/codex/tree/nodes"
)

type Accessor func(string) *nodes.Attribute

func (accessor Accessor) Relation() *nodes.Relation {
  return accessor("").Relation
}

func (accessor Accessor) Name() string {
  return accessor("").Relation.Name
}

func (accessor Accessor) Project(projections ...interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Project(projections...)
}

func (accessor Accessor) Where(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Where(expr)
}

func (accessor Accessor) InnerJoin(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).InnerJoin(expr)
}

func (accessor Accessor) OuterJoin(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).OuterJoin(expr)
}

func (accessor Accessor) From(relation *nodes.Relation) *SelectManager {
  mgmt := new(SelectManager)
  mgmt.Tree = factory.SelectStatement(relation)
  mgmt.Context = mgmt.Tree.Cores[0]
  return mgmt
}

func (accessor Accessor) Insert(expr ...interface{}) *InsertManager {
  stmt := factory.InsertStatement(accessor.Relation())
  mgmt := new(InsertManager)
  mgmt.Tree = stmt
  return mgmt.Insert(expr...)
}

func (accessor Accessor) Set(expr ...interface{}) *UpdateManager {
  stmt := factory.UpdateStatement(accessor.Relation())
  mgmt := new(UpdateManager)
  mgmt.Tree = stmt
  return mgmt.Set(expr...)
}

func (accessor Accessor) Delete(expr interface{}) *DeleteManager {
  stmt := factory.DeleteStatement(accessor.Relation())
  mgmt := new(DeleteManager)
  mgmt.Tree = stmt
  return mgmt.Delete(expr)
}

func (accessor Accessor) ToSql() string {
  return accessor.From(accessor.Relation()).ToSql()
}
