package librarian

import (
  "fmt"
  "librarian/nodes"
)

type UpdateManager struct {
  Tree     *nodes.UpdateStatementNode
  Context  *nodes.UpdateStatementNode
  Relation *nodes.RelationNode
  Engine   interface{}
}

func (mgmt *UpdateManager) Limit(take interface{}) *UpdateManager {
  mgmt.Context.Limit = take
  return mgmt
}

func (mgmt *UpdateManager) Where(comparison interface{}) *UpdateManager {
  mgmt.Context.Wheres = append(mgmt.Context.Wheres, comparison)
  return mgmt
}

func (mgmt *UpdateManager) Set(assignment interface{}) *UpdateManager {
  mgmt.Context.Values = append(mgmt.Context.Values, assignment)
  return mgmt
}

func (mgmt *UpdateManager) ToSql() string {
  var engine string
  if 0 == len(mgmt.Relation.Engine) {
    engine = "to_sql"
  }
  visitor, ok := VISITORS[engine]
  if !ok {
    panic(fmt.Sprintf("No engine found for %v.\n", engine))
  }
  return visitor.Accept(mgmt.Tree)
}

func NewUpdateManager(relation *nodes.RelationNode) *UpdateManager {
  tree := nodes.UpdateStatement(relation)
  return &UpdateManager{tree, tree, relation, nil}
}
