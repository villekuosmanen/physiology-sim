package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Muscle struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Muscle)(nil)
var _ control.Controller = (*Muscle)(nil)

func ConstructMuscle(consumer circulation.BloodConsumer) *Muscle {
	return &Muscle{
		vascularity: NewVascularity(VascularityRating4, &metabolism.OxygenMetaboliser{}),
		consumer:    consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (b *Muscle) AcceptBlood(bl circulation.Blood) {
	b.vascularity.AcceptBlood(bl)
}

// Act implements control.Controller
func (b *Muscle) Act() {
	// Metabolise
	bl := b.vascularity.Process()
	b.consumer.AcceptBlood(bl)
}

func (b *Muscle) BloodQuantity() float64 {
	return b.vascularity.BloodQuantity()
}
