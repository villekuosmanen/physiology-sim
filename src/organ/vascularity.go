package organ

import (
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
	"github.com/villekuosmanen/physiology-sim/src/systems/nerve"
)

type Vascularity struct {
	Gates                 []*OrganGate
	Metaboliser           metabolism.Metaboliser
	gateMovingRate        float64
	snsSignalHandleMethod nerve.SNSSignalHandleMethod
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

func NewVascularity(
	numberOfGates int,
	metaboliser metabolism.Metaboliser,
	snsSignalHandleMethod nerve.SNSSignalHandleMethod,
) *Vascularity {
	gates := []*OrganGate{}
	for i := 0; i < numberOfGates; i++ {
		gates = append(gates, &OrganGate{
			blood: &circulation.Blood{},
		})
	}
	return &Vascularity{
		Gates:                 gates,
		Metaboliser:           metaboliser,
		gateMovingRate:        circulation.EmptyRateVascularity,
		snsSignalHandleMethod: snsSignalHandleMethod,
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
		norepinephrineEffect := 1.0
		// if v.snsSignalHandleMethod == nerve.SNSSignalHandleMethodExpand {
		// 	norepinephrineEffect = 1.5 - (0.75 * gate.blood.Norepinephrine)
		// } else if v.snsSignalHandleMethod == nerve.SNSSignalHandleMethodContract {
		// 	norepinephrineEffect = 0.75 + (gate.blood.Norepinephrine * 0.75)
		// }
		temp = gate.blood.Extract(v.gateMovingRate * norepinephrineEffect)
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

func (v *Vascularity) LacticAcid() float64 {
	return v.Gates[len(v.Gates)-1].blood.LacticAcid
}

func (v *Vascularity) Norepinephrine() float64 {
	return v.Gates[len(v.Gates)-1].blood.Norepinephrine
}
