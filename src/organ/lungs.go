package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Lungs struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Lungs)(nil)
var _ control.MonitorableController = (*Lungs)(nil)

func ConstructLungs(consumer circulation.BloodConsumer) *Lungs {
	return &Lungs{
		vascularity: NewVascularity(VascularityRating3, &metabolism.OxygenMetaboliser{}),
		consumer:    consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Lungs) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements control.Controller
func (b *Lungs) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

// Monitor implements control.Controller
func (b *Lungs) Monitor() *control.BloodStatistics {
	return &control.BloodStatistics{
		ComponentName: "Lungs",
		BloodQuantity: b.vascularity.BloodQuantity(),
	}
}
