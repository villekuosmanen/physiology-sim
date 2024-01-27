package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Kidney struct {
	// contains a reservoir for blood
	name      string
	blood     *circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Kidney)(nil)
var _ control.MonitorableController = (*Kidney)(nil)

func ConstructKidney(name string, consumer circulation.BloodConsumer) *Kidney {
	return &Kidney{
		name:      name,
		blood:     &circulation.Blood{},
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

// Monitor implements control.Controller
func (v *Kidney) Monitor() *control.BloodStatistics {
	return &control.BloodStatistics{
		ComponentName: v.name,
		BloodQuantity: v.blood.Quantity,
	}
}
