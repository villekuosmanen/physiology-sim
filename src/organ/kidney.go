package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
	"github.com/villekuosmanen/physiology-sim/src/systems/nerve"
)

type Kidney struct {
	// contains a reservoir for blood
	name        string
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Kidney)(nil)
var _ simulation.MonitorableController = (*Kidney)(nil)

// NOTE:
// kidneys are not very vascular but have a massive blood supply

func ConstructKidney(name string, consumer circulation.BloodConsumer) *Kidney {
	return &Kidney{
		name: name,
		vascularity: NewVascularity(
			VascularityRating3,
			&metabolism.OxygenConsumer{},
			nerve.SNSSignalHandleMethodContract,
		),
		consumer: consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Kidney) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements simulation.Controller
func (b *Kidney) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

// Monitor implements simulation.Controller
func (v *Kidney) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName:       v.name,
		BloodQuantity:       v.vascularity.BloodQuantity(),
		HasOxygenSaturation: true,
		OxygenSaturation:    v.vascularity.OxygenSaturation(),
	}
}
