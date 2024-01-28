package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
	"github.com/villekuosmanen/physiology-sim/src/systems/nerve"
)

type Muscle struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	metaboliser *metabolism.MuscleMetaboliser
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Muscle)(nil)
var _ simulation.Controller = (*Muscle)(nil)

func ConstructMuscle(consumer circulation.BloodConsumer) *Muscle {
	metaboliser := metabolism.NewMuscleMetaboliser()
	return &Muscle{
		metaboliser: metaboliser,
		vascularity: NewVascularity(
			VascularityRating4,
			metaboliser,
			nerve.SNSSignalHandleMethodExpand,
		),
		consumer: consumer,
	}
}

func (b *Muscle) SetMetabolicRate(new metabolism.MET) {
	b.metaboliser.SetMetabolicRate(new)
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Muscle) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements simulation.Controller
func (b *Muscle) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

func (b *Muscle) BloodQuantity() float64 {
	return b.vascularity.BloodQuantity()
}
