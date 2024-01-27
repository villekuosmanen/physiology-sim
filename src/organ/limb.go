package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Limb struct {
	muscle      *Muscle
	fat         *Fat
	muscleShare float64
}

var _ circulation.BloodConsumer = (*Limb)(nil)
var _ control.Controller = (*Limb)(nil)

func ConstructLimb(muscleShare float64, consumer circulation.BloodConsumer) *Limb {
	return &Limb{
		muscle:      ConstructMuscle(consumer),
		fat:         ConstructFat(consumer),
		muscleShare: muscleShare,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Limb) AcceptBlood(bl circulation.Blood) {
	// assume muscle takes up twice as much blood as fat
	muscleBloodShare := (b.muscleShare * 2) / ((b.muscleShare * 2) + (1 - b.muscleShare))

	muscleBlood := bl.Extract(muscleBloodShare)
	b.muscle.AcceptBlood(muscleBlood)
	b.fat.AcceptBlood(bl)
}

// Act implements control.Controller
func (b *Limb) Act() {
	// Limb asks for its constituents to act.
	b.muscle.Act()
	b.fat.Act()
}
