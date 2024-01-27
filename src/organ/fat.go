package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Fat struct {
	// contains a reservoir for blood
	vascularity *Vascularity
	consumer    circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Fat)(nil)
var _ control.Controller = (*Fat)(nil)

func ConstructFat(consumer circulation.BloodConsumer) *Fat {
	return &Fat{
		vascularity: NewVascularity(VascularityRating1, &metabolism.OxygenConsumer{}),
		consumer:    consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (f *Fat) AcceptBlood(bl circulation.Blood) {
	f.vascularity.AcceptBlood(bl)
}

// Act implements control.Controller
func (f *Fat) Act() {
	// Metabolise
	bl := f.vascularity.Process()
	f.consumer.AcceptBlood(bl)
}

func (f *Fat) BloodQuantity() float64 {
	return f.vascularity.BloodQuantity()
}
