package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Lungs struct {
	// contains a reservoir for blood
	blood     *circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Lungs)(nil)
var _ control.MonitorableController = (*Lungs)(nil)

func ConstructLungs(consumer circulation.BloodConsumer) *Lungs {
	return &Lungs{
		blood:     &circulation.Blood{},
		emptyRate: circulation.EmptyRateAverage,
		consumer:  consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Lungs) AcceptBlood(bl circulation.Blood) {
	b.blood.Merge(bl)
}

// Act implements control.Controller
func (b *Lungs) Act() {
	// Currently the Lungs does nothing useful.

	// move blood away from the Lungs
	bl := b.blood.Extract(b.emptyRate)
	b.consumer.AcceptBlood(bl)
}

// Monitor implements control.Controller
func (b *Lungs) Monitor() *control.BloodStatistics {
	return &control.BloodStatistics{
		ComponentName: "Lungs",
		BloodQuantity: b.blood.Quantity,
	}
}
