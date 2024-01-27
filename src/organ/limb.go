package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
)

type Limb struct {
	Muscle *Muscle

	name        string
	fat         *Fat
	muscleShare float64
}

var _ circulation.BloodConsumer = (*Limb)(nil)
var _ simulation.MonitorableController = (*Limb)(nil)

func ConstructLimb(name string, muscleShare float64, consumer circulation.BloodConsumer) *Limb {
	return &Limb{
		name:        name,
		Muscle:      ConstructMuscle(consumer),
		fat:         ConstructFat(consumer),
		muscleShare: muscleShare,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Limb) AcceptBlood(bl circulation.Blood) {
	// assume muscle takes up twice as much blood as fat
	muscleBloodShare := (b.muscleShare * 2) / ((b.muscleShare * 2) + (1 - b.muscleShare))

	muscleBlood := bl.Extract(muscleBloodShare)
	b.Muscle.AcceptBlood(muscleBlood)
	b.fat.AcceptBlood(bl)
}

// Act implements simulation.Controller
func (b *Limb) Act() {
	// Limb asks for its constituents to act.
	b.Muscle.Act()
	b.fat.Act()
}

// Monitor implements simulation.Controller
func (b *Limb) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName: b.name,
		BloodQuantity: b.Muscle.BloodQuantity() + b.fat.BloodQuantity(),
	}
}
