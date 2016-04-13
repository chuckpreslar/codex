package interfaces

// Acceptor ...
type Acceptor interface {
	Accept(Visitor) string
}
