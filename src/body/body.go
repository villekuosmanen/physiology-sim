package body

import (
	"time"

	"github.com/villekuosmanen/physiology-sim/src/organ"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
)

// TODO this should not be a constant
const heartRate = 80

type Body struct {
	Heart            organ.Heart
	Aorta            circulation.Vessel
	SuperiorVenaCava circulation.Vessel
	InferiorVenaCava circulation.Vessel

	PulmonaryArtery circulation.Vessel
	PulmonaryVein   circulation.Vessel
	Lungs           organ.Lungs

	Brain       organ.Brain
	Liver       organ.Liver
	LeftKidney  organ.Kidney
	RightKidney organ.Kidney

	// To be added
	// digestive tract
	// Portal Vein

	LeftBreast  organ.TorsoPart
	RightBreast organ.TorsoPart
	Abdomen     organ.TorsoPart

	RightArm organ.Limb
	LeftArm  organ.Limb
	RightLeg organ.Limb
	LeftLeg  organ.Limb
}

func (b *Body) Run(frequency float64, realtime bool) {
	// run forever in given Hz
	untilNextHeartbeat := 0.0

	var t *time.Ticker
	if realtime {
		t = time.NewTicker(time.Second / time.Duration(frequency))
	} else {
		// 100 times faster than otherwise
		t = time.NewTicker(time.Second / (time.Duration(frequency) * 100))
	}

	for {
		// wait for a tick
		<-t.C

		if untilNextHeartbeat <= 0 {
			b.Heart.Beat()
			untilNextHeartbeat = heartRate * frequency
		} else {
			untilNextHeartbeat -= 1
		}

		b.Act()
	}
}

func (b *Body) Act() {
	b.Heart.Myocardium.Act()
	b.Aorta.Act()

	b.PulmonaryArtery.Act()
	b.Lungs.Act()
	b.InferiorVenaCava.Act()

	b.Brain.Act()
	b.Liver.Act()
	b.LeftKidney.Act()
	b.RightKidney.Act()

	b.LeftBreast.Act()
	b.RightBreast.Act()
	b.Abdomen.Act()

	b.RightArm.Act()
	b.LeftArm.Act()
	b.RightLeg.Act()
	b.LeftLeg.Act()

	b.SuperiorVenaCava.Act()
	b.InferiorVenaCava.Act()
}
