package entity

import (
	"fmt"
	"github.com/accnameowl/govect"
)

// Vector ...
type Vector govect.Vector

// Entity ...
type Entity struct {
	Position Vector
	Collition
}
