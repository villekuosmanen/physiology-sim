package circulation

import "github.com/villekuosmanen/physiology-sim/src/systems/control"

type Vessel struct {
	// contains a reservoir for blood
	blood     Blood
	emptyRate float64 // the rate at which the vessel empties
	consumers []BloodConsumer
}

var _ BloodConsumer = (*Vessel)(nil)
var _ control.Controller = (*Vessel)(nil)

func ConstructVessel(consumers []BloodConsumer, isArtery bool) *Vessel {
	emptyRate := EmptyRateFast
	if isArtery {
		emptyRate = EmptyRateVeryFast
	}

	return &Vessel{
		blood:     Blood{},
		emptyRate: emptyRate,
		consumers: consumers,
	}
}

// Act implements control.Controller
func (v *Vessel) Act() {
	// At each tick, a share of the blood avaiable in the artery is sent to its outflows
	// TODO: this is unrealistic.
	// Each of the consumers needs to order a specific share of blood being consumed from the vessel.
	// This ensures bigger organs receive more blood than smaller ones.
	allBlood := v.blood.Extract(v.emptyRate)
	bloodPerConsumer := DivideBlood(allBlood, len(v.consumers))

	for _, c := range v.consumers {
		c.AcceptBlood(bloodPerConsumer)
	}
}

// AcceptBlood implements BloodConsumer
func (v *Vessel) AcceptBlood(b Blood) {
	v.blood.Merge(b)
}
