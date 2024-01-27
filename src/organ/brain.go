package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Brain struct {
	// contains a reservoir for blood
	blood     circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Brain)(nil)
var _ control.Controller = (*Brain)(nil)

func ConstructBrain(consumer circulation.BloodConsumer) *Liver {
	return &Liver{
		blood:     circulation.Blood{},
		emptyRate: circulation.EmptyRateSlow,
		consumer:  consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Brain) AcceptBlood(bl circulation.Blood) {
	b.blood.Merge(bl)
}

// Act implements control.Controller
func (b *Brain) Act() {
	// Currently the brain does nothing useful.

	// move blood away from the brain
	bl := b.blood.Extract(b.emptyRate)
	b.consumer.AcceptBlood(bl)
}
