package tracker

import (
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrNoRecords = fmt.Errorf("no records")
var ErrIDAlreadyExists = fmt.Errorf("id already exist")

func ErrPositionNotFound(position int) error {
	return fmt.Errorf("element number %d not found", position)
}
