package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
)

type TorsoPart struct {
	Muscle *Muscle

	name        string
	fat         *Fat
	muscleShare float64
}

var _ circulation.BloodConsumer = (*TorsoPart)(nil)
var _ simulation.MonitorableController = (*TorsoPart)(nil)

func ConstructTorsoPart(name string, muscleShare float64, consumer circulation.BloodConsumer) *TorsoPart {
	return &TorsoPart{
		name:        name,
		Muscle:      ConstructMuscle(consumer),
		fat:         ConstructFat(consumer),
		muscleShare: muscleShare,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *TorsoPart) AcceptBlood(bl circulation.Blood) {
	// assume muscle takes up twice as much blood as fat
	muscleBloodShare := (b.muscleShare * 2) / ((b.muscleShare * 2) + (1 - b.muscleShare))

	muscleBlood := bl.Extract(muscleBloodShare)
	b.Muscle.AcceptBlood(muscleBlood)
	b.fat.AcceptBlood(bl)
}

// Act implements simulation.Controller
func (b *TorsoPart) Act() {
	// TorsoPart asks for its constituents to act.
	b.Muscle.Act()
	b.fat.Act()
}

// Monitor implements simulation.Controller
func (b *TorsoPart) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName: b.name,
		BloodQuantity: b.Muscle.BloodQuantity() + b.fat.BloodQuantity(),
	}
}
