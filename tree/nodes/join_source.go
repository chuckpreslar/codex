package nodes

// JoinSource is a specific Binary node.
type JoinSource struct {
  Left  *Relation     // Left child of the JoinSource node, a pointer to a Relation.
  Right []interface{} // Right child of the JoinSource node contains joins and their instructions
}
