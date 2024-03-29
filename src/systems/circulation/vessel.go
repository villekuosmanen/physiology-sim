package circulation

import (
	"github.com/villekuosmanen/physiology-sim/src/simulation"
)

type Vessel struct {
	// contains a reservoir for blood
	name      string
	blood     *Blood
	emptyRate float64 // the rate at which the vessel empties
	consumers []ConsumerWithBloodSupply
}

type ConsumerWithBloodSupply struct {
	Consumer    BloodConsumer
	BloodSupply float64
}

const (
	VesselSizeHuge   = 0.75
	VesselSizeLarge  = 1
	VesselSizeMedium = 1.5
)

var _ BloodConsumer = (*Vessel)(nil)
var _ simulation.MonitorableController = (*Vessel)(nil)

func ConstructVessel(
	name string,
	vesselSize float64,
	consumers []ConsumerWithBloodSupply,
	isArtery bool,
) *Vessel {
	emptyRate := EmptyRateVein
	if isArtery {
		emptyRate = EmptyRateArtery
	}

	emptyRate *= vesselSize

	return &Vessel{
		name:      name,
		blood:     &Blood{},
		emptyRate: emptyRate,
		consumers: consumers,
	}
}

// Act implements simulation.Controller
func (v *Vessel) Act() {
	// At each tick, a share of the blood avaiable in the artery is sent to its outflows
	// This represents how fast blood circulates in the vessels.
	// Autonomous nervous system affects it through neurotransmitters.

	norepinephrineEffect := 0.75 + (v.blood.Norepinephrine * 0.75)
	effectiveEmptyRate := v.emptyRate * norepinephrineEffect

	allBlood := v.blood.Extract(effectiveEmptyRate)
	totalWeight := 0.0
	for _, consumer := range v.consumers {
		totalWeight += consumer.BloodSupply
	}

	for _, consumer := range v.consumers {
		bloodPerConsumer := RemoveFrom(allBlood, (consumer.BloodSupply / totalWeight))
		consumer.Consumer.AcceptBlood(bloodPerConsumer)
	}
}

// Monitor implements simulation.Controller
func (v *Vessel) Monitor() *simulation.BloodStatistics {
	return &simulation.BloodStatistics{
		ComponentName: v.name,
		BloodQuantity: v.blood.Quantity,
	}
}

// AcceptBlood implements BloodConsumer
func (v *Vessel) AcceptBlood(b Blood) {
	v.blood.Merge(b)
}
