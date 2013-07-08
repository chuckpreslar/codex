package nodes

// JoinSource is a specific Binary node.
type JoinSource struct {
  Left  interface{}   // Left child of the JoinSource node, usually a pointer to a Relation.
  Right []interface{} // Right child of the JoinSource node contains joins and their instructions
}
