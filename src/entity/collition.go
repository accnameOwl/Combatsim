package entity

import "github.com/accnameowl/govect"

// Collition ...
type Collition govect.Vector

// Collition calc consts
const(
	CollitionCrossX = 1
	CollitionCrossY = 1
	CollitionCrossedY = 1
	CollitionCrossedX = 1
)

// Cross ...
func (c *Collition) Cross(p *Collition) bool {
	// cross, err := govect.Cross(c, p)
}
// Crossed ...
func (c *Collition) Crossed(p *Collition) bool {
	// cross, err := govect.Cross(c, p)
}

// Enter ...
func (c *Collition) Enter() bool {

}

// Entered ...
func (c *Collition) Entered() bool {

}

// Exit ...
func (c *Collition) Exit() bool {

}

// Exited ...
func (c *Collition) Exited() bool {

}
