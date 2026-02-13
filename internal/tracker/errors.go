package tracker

import (
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrIDAlreadyExists = fmt.Errorf("id already exist")
