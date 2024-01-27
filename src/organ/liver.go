package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Liver struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Liver)(nil)
var _ simulation.MonitorableController = (*Liver)(nil)

func ConstructLiver(consumer circulation.BloodConsumer) *Liver {
	return &Liver{
		vascularity: NewVascularity(VascularityRating8, &metabolism.OxygenConsumer{}),
		consumer:    consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Liver) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements simulation.Controller
func (b *Liver) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

// Monitor implements simulation.Controller
func (b *Liver) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName: "Liver",
		BloodQuantity: b.vascularity.BloodQuantity(),
	}
}
