package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Muscle struct {
	// contains a reservoir for blood
	blood     circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Muscle)(nil)
var _ control.Controller = (*Muscle)(nil)

func ConstructMuscle(consumer circulation.BloodConsumer) *Muscle {
	return &Muscle{
		blood:     circulation.Blood{},
		emptyRate: circulation.EmptyRateAverage,
		consumer:  consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Muscle) AcceptBlood(bl circulation.Blood) {
	b.blood.Merge(bl)
}

// Act implements control.Controller
func (b *Muscle) Act() {
	// Currently the Muscle does nothing useful.

	// move blood away from the Muscle
	bl := b.blood.Extract(b.emptyRate)
	b.consumer.AcceptBlood(bl)
}
