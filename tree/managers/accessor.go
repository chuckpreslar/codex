package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// Accessor is a type def of a function that returns an AttributeNode
type Accessor func(interface{}) *nodes.AttributeNode

// Returns the RelationNode scoped to the Accessor
func (accessor Accessor) Relation() *nodes.RelationNode {
  return accessor("").Relation
}

// Returns the name of the RelationNode scoped to the Accessor
func (accessor Accessor) Name() string {
  return accessor("").Relation.Name
}

// Returns a pointer to a SelectManager with the initial projections provided.
func (accessor Accessor) Project(projections ...interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Project(projections...)
}

// Returns a pointer to a SelectManager with the initial filter provided.
func (accessor Accessor) Where(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Where(expr)
}

// Returns a pointer to a SelectManager with an initial InnerJoinNode.
func (accessor Accessor) InnerJoin(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).InnerJoin(expr)
}

// Returns a pointer to a SelectManager with an initial OuterJoinNode.
func (accessor Accessor) OuterJoin(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).OuterJoin(expr)
}

// Returns a pointer to a SelectManager with an initial Ordering.
func (accessor Accessor) Order(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Order(expr)
}

// Returns a pointer to a SelectManager with for the given RelationNode.
func (accessor Accessor) From(relation *nodes.RelationNode) *SelectManager {
  return Selection(relation)
}

// Returns a pointer to a InsertManager with initial the values provided.
func (accessor Accessor) Insert(expr ...interface{}) *InsertManager {
  return Insertion(accessor.Relation()).Insert(expr...)
}

// Returns a pointer to a UpdateManager with initial the columns provided.
func (accessor Accessor) Set(expr ...interface{}) *UpdateManager {
  return Modification(accessor.Relation()).Set(expr...)
}

// Returns a pointer to a DeleteManager with the initial filter provided.
func (accessor Accessor) Delete(expr interface{}) *DeleteManager {
  return Deletion(accessor.Relation()).Delete(expr)
}

// Returns string, error generated by calling ToSql on an SelectManager.
func (accessor Accessor) ToSql() (string, error) {
  return accessor.From(accessor.Relation()).ToSql()
}
