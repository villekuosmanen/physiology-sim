package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Kidney struct {
	// contains a reservoir for blood
	name        string
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Kidney)(nil)
var _ control.MonitorableController = (*Kidney)(nil)

// NOTE:
// kidneys are not very vascular but have a massive blood supply

func ConstructKidney(name string, consumer circulation.BloodConsumer) *Kidney {
	return &Kidney{
		name:        name,
		vascularity: NewVascularity(VascularityRating4, &metabolism.OxygenMetaboliser{}),
		consumer:    consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Kidney) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements control.Controller
func (b *Kidney) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

// Monitor implements control.Controller
func (v *Kidney) Monitor() *control.BloodStatistics {
	return &control.BloodStatistics{
		ComponentName: v.name,
		BloodQuantity: v.vascularity.BloodQuantity(),
	}
}
