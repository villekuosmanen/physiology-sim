package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Liver struct {
	// contains a reservoir for blood
	blood     circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Liver)(nil)
var _ control.Controller = (*Liver)(nil)

func ConstructLiver(consumer circulation.BloodConsumer) *Liver {
	return &Liver{
		blood:     circulation.Blood{},
		emptyRate: circulation.EmptyRateVerySlow,
		consumer:  consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Liver) AcceptBlood(bl circulation.Blood) {
	b.blood.Merge(bl)
}

// Act implements control.Controller
func (b *Liver) Act() {
	// Currently the Liver does nothing useful.

	// move blood away from the Liver
	bl := b.blood.Extract(b.emptyRate)
	b.consumer.AcceptBlood(bl)
}
