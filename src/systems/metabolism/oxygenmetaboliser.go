package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type OxygenMetaboliser struct{}

var _ Metaboliser = (*OxygenMetaboliser)(nil)

// Metabolise implements Metaboliser.
func (*OxygenMetaboliser) Metabolise(b *circulation.Blood) {
	// no-op for now
}
