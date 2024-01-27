package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type Vascularity struct {
	Gates          []*OrganGate
	Metaboliser    metabolism.Metaboliser
	gateMovingRate float64
}

type OrganGate struct {
	blood *circulation.Blood
}

const (
	VascularityRating10 = 10
	VascularityRating8  = 8
	VascularityRating6  = 6
	VascularityRating5  = 5
	VascularityRating4  = 4
	VascularityRating3  = 3
	VascularityRating2  = 2
	VascularityRating1  = 1
)

func NewVascularity(numberOfGates int, metaboliser metabolism.Metaboliser) *Vascularity {
	gates := []*OrganGate{}
	for i := 0; i < numberOfGates; i++ {
		gates = append(gates, &OrganGate{
			blood: &circulation.Blood{},
		})
	}
	return &Vascularity{
		Gates:          gates,
		Metaboliser:    metaboliser,
		gateMovingRate: circulation.EmptyRateVascularity,
	}
}

// AcceptBlood implements circulation.BloodConsumer
func (v *Vascularity) AcceptBlood(bl circulation.Blood) {
	v.Gates[0].blood.Merge(bl)
}

func (v *Vascularity) Process() circulation.Blood {
	var temp circulation.Blood
	for _, gate := range v.Gates {
		// this is a no-op at the first gate
		gate.blood.Merge(temp)

		// metabolise blood
		v.Metaboliser.Metabolise(gate.blood)

		// set some of current blood into the next gate
		temp = gate.blood.Extract(v.gateMovingRate)
	}

	return temp
}

func (v *Vascularity) BloodQuantity() float64 {
	qty := 0.0
	for _, gate := range v.Gates {
		qty += gate.blood.Quantity
	}

	return qty
}

func (v *Vascularity) OxygenSaturation() float64 {
	return v.Gates[len(v.Gates)-1].blood.OxygenSaturation
}
