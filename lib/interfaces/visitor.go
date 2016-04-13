package interfaces

// Visitor ...
type Visitor interface {
	Visit(Acceptor) string
}
