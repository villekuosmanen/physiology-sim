package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
	"github.com/villekuosmanen/physiology-sim/src/systems/nerve"
)

type Lungs struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Lungs)(nil)
var _ simulation.MonitorableController = (*Lungs)(nil)

func ConstructLungs(consumer circulation.BloodConsumer) *Lungs {
	return &Lungs{
		vascularity: NewVascularity(
			VascularityRating3,
			metabolism.NewLungMetaboliser(),
			nerve.SNSSignalHandleMethodNothing,
		),
		consumer: consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Lungs) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements simulation.Controller
func (b *Lungs) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

// Monitor implements simulation.Controller
func (b *Lungs) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName: "Lungs",
		BloodQuantity: b.vascularity.BloodQuantity(),
	}
}
