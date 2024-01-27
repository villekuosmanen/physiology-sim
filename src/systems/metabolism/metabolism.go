package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type Metaboliser interface {
	Metabolise(b *circulation.Blood)
}

// scratch pad
// - each organ has metabolism gates that work exactly the same except two things vary:
//   - lenght of gates
//   - the metabolism function
