package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type TorsoPart struct {
	name        string
	muscle      *Muscle
	fat         *Fat
	muscleShare float64
}

var _ circulation.BloodConsumer = (*TorsoPart)(nil)
var _ control.MonitorableController = (*TorsoPart)(nil)

func ConstructTorsoPart(name string, muscleShare float64, consumer circulation.BloodConsumer) *TorsoPart {
	return &TorsoPart{
		name:        name,
		muscle:      ConstructMuscle(consumer),
		fat:         ConstructFat(consumer),
		muscleShare: muscleShare,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *TorsoPart) AcceptBlood(bl circulation.Blood) {
	// assume muscle takes up twice as much blood as fat
	muscleBloodShare := (b.muscleShare * 2) / ((b.muscleShare * 2) + (1 - b.muscleShare))

	muscleBlood := bl.Extract(muscleBloodShare)
	b.muscle.AcceptBlood(muscleBlood)
	b.fat.AcceptBlood(bl)
}

// Act implements control.Controller
func (b *TorsoPart) Act() {
	// TorsoPart asks for its constituents to act.
	b.muscle.Act()
	b.fat.Act()
}

// Monitor implements control.Controller
func (b *TorsoPart) Monitor() *control.BloodStatistics {
	return &control.BloodStatistics{
		ComponentName: b.name,
		BloodQuantity: b.muscle.BloodQuantity() + b.fat.BloodQuantity(),
	}
}
