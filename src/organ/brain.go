package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
	"github.com/villekuosmanen/physiology-sim/src/systems/nerve"
)

type Brain struct {
	// contains a reservoir for blood
	vascularity    *Vascularity
	brainRegulator *nerve.BrainRegulator
	consumer       circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Brain)(nil)
var _ simulation.MonitorableController = (*Brain)(nil)

func ConstructBrain(consumer circulation.BloodConsumer) *Brain {
	return &Brain{
		vascularity:    NewVascularity(VascularityRating8, &metabolism.OxygenConsumer{}),
		brainRegulator: &nerve.BrainRegulator{},
		consumer:       consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Brain) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

func (b *Brain) Act() {
	// Metabolise
	bl := b.vascularity.Process()

	// Regulate neurotransmitters
	bl = b.brainRegulator.Regulate(bl)

	b.consumer.AcceptBlood(bl)
}

// Monitor implements simulation.Controller
func (b *Brain) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName:       "Brain",
		BloodQuantity:       b.vascularity.BloodQuantity(),
		HasOxygenSaturation: true,
		OxygenSaturation:    b.vascularity.OxygenSaturation(),
		HasLacticAcid:       true,
		LacticAcid:          b.vascularity.LacticAcid(),
		HasNorepinephrine:   true,
		Norepinephrine:      b.vascularity.Norepinephrine(),
	}
}
