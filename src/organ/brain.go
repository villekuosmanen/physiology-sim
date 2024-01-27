package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Brain struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Brain)(nil)
var _ control.MonitorableController = (*Brain)(nil)

func ConstructBrain(consumer circulation.BloodConsumer) *Brain {
	return &Brain{
		vascularity: NewVascularity(VascularityRating8, &metabolism.OxygenConsumer{}),
		consumer:    consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Brain) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements control.Controller
func (b *Brain) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

// Monitor implements control.Controller
func (b *Brain) Monitor() *control.BloodStatistics {
	return &control.BloodStatistics{
		ComponentName:       "Brain",
		BloodQuantity:       b.vascularity.BloodQuantity(),
		HasOxygenSaturation: true,
		OxygenSaturation:    b.vascularity.OxygenSaturation(),
	}
}
