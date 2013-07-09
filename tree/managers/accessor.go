package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type Accessor func(interface{}) *nodes.AttributeNode

func (accessor Accessor) Relation() *nodes.RelationNode {
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

func (accessor Accessor) From(relation *nodes.RelationNode) *SelectManager {
  return Selection(relation)
}

func (accessor Accessor) Insert(expr ...interface{}) *InsertManager {
  return Insertion(accessor.Relation()).Insert(expr...)
}

func (accessor Accessor) Set(expr ...interface{}) *UpdateManager {
  return Modification(accessor.Relation()).Set(expr...)
}

func (accessor Accessor) Delete(expr interface{}) *DeleteManager {
  return Deletion(accessor.Relation()).Delete(expr)
}

func (accessor Accessor) ToSql() (string, error) {
  return accessor.From(accessor.Relation()).ToSql()
}
