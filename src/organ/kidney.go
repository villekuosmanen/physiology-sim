package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Kidney struct {
	// contains a reservoir for blood
	blood     circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Kidney)(nil)
var _ control.Controller = (*Kidney)(nil)

func ConstructKidney(consumer circulation.BloodConsumer) *Kidney {
	return &Kidney{
		blood:     circulation.Blood{},
		emptyRate: circulation.EmptyRateSlow,
		consumer:  consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Kidney) AcceptBlood(bl circulation.Blood) {
	b.blood.Merge(bl)
}

// Act implements control.Controller
func (b *Kidney) Act() {
	// Currently the Kidney does nothing useful.

	// move blood away from the Kidney
	bl := b.blood.Extract(b.emptyRate)
	b.consumer.AcceptBlood(bl)
}
