package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/control"
)

type Fat struct {
	// contains a reservoir for blood
	blood     *circulation.Blood
	emptyRate float64 // the rate at which the vessel empties
	consumer  circulation.BloodConsumer
}

var _ circulation.BloodConsumer = (*Fat)(nil)
var _ control.Controller = (*Fat)(nil)

func ConstructFat(consumer circulation.BloodConsumer) *Fat {
	return &Fat{
		blood:     &circulation.Blood{},
		emptyRate: circulation.EmptyRateAverage,
		consumer:  consumer,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (f *Fat) AcceptBlood(bl circulation.Blood) {
	f.blood.Merge(bl)
}

// Act implements control.Controller
func (f *Fat) Act() {
	// Currently the Fat does nothing useful.

	// move blood away from the Fat
	bl := f.blood.Extract(f.emptyRate)
	f.consumer.AcceptBlood(bl)
}
