// Package nodes provides nodes to use in codex AST's.
package nodes

// SelectCoreNode is a Nary node, normally contained in a SelectStatement node.
type SelectCoreNode struct {
	Relation    *RelationNode   // Pointer to the relation the SelectCore is acting on.
	Source      *JoinSourceNode // JoinSouce for joining other SQL tables.
	Projections []interface{}   // Projections is an array, normally columns found on the SQL table.
	Wheres      []interface{}   // Wheres is an array of filters for the acting on the SelectCore.
	Groups      []interface{}   // GROUP BY nodes.
	Having      interface{}     // HAVING expression.
}

// SelectCoreNode factory method.
func SelectCore(relation *RelationNode) (core *SelectCoreNode) {
	core = new(SelectCoreNode)
	core.Source = JoinSource(relation)
	core.Relation = relation
	core.Wheres = make([]interface{}, 0)
	core.Projections = make([]interface{}, 0)
	core.Groups = make([]interface{}, 0)
	return
}
