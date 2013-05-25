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

var BadSessionError LibrarianError = LibrarianError{"Bad or no database session was provided.", time.Now()}
var BadArgsError LibrarianError = LibrarianError{"Bad arguments supplied.", time.Now()}
var BadResultsError LibrarianError = LibrarianError{"Unexpected results.", time.Now()}
