package librarian

import (
	"fmt"
	"time"
)

type LibrarianError struct {
	details   string
	timestamp time.Time
}

func (l LibrarianError) Error() string {
	return fmt.Sprintf("Librarian : %v : %s", l.timestamp, l.details)
}

var NoSessionError LibrarianError = LibrarianError{"No database session was provided.", time.Now()}
