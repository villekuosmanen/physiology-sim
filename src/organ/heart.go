package organ

import (
	"math"

	"github.com/villekuosmanen/physiology-sim/src/simulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
)

type Heart struct {
	// contains a muscle
	Myocardium Muscle

	// contains a pacemaker
	pacemaker pacemaker

	// contains two atria
	// contains two ventricles
	RightAtrium    atrium    // from systemic
	LeftAtrium     atrium    // from pulmonary
	rightVentricle ventricle // to pulmonary
	leftVentricle  ventricle // to systemic
}

func ConstructHeart() *Heart {
	leftAtrium := atrium{
		Blood: &circulation.Blood{},
	}
	rightAtrium := atrium{
		Blood: &circulation.Blood{},
	}

	myocardium := ConstructMuscle(&rightAtrium)

	leftVentricle := ventricle{
		// artery to be set below
		Blood: &circulation.Blood{
			Quantity:         50,
			OxygenSaturation: 0.9,
		},
	}
	rightVentricle := ventricle{
		// artery to be set below
		Blood: &circulation.Blood{
			Quantity:         50,
			OxygenSaturation: 0.9,
		},
	}

	return &Heart{
		Myocardium:     *myocardium,
		pacemaker:      pacemaker{},
		RightAtrium:    rightAtrium,
		LeftAtrium:     leftAtrium,
		rightVentricle: rightVentricle,
		leftVentricle:  leftVentricle,
	}
}

func (h *Heart) SetConsumers(aorta, pulmonaryArtery circulation.BloodConsumer) {
	h.leftVentricle.Artery = aorta
	h.rightVentricle.Artery = pulmonaryArtery
}

type atrium struct {
	Blood *circulation.Blood
}

var _ circulation.BloodConsumer = (*atrium)(nil)

func (a *atrium) AcceptBlood(b circulation.Blood) {
	a.Blood.Merge(b)
}

type ventricle struct {
	Blood  *circulation.Blood
	Artery circulation.BloodConsumer
}

type pacemaker struct {
}

func (h *Heart) Beat() float64 {
	// blood moves from atria to ventricles
	br := h.RightAtrium.Blood.Extract(0.95)
	h.rightVentricle.Blood.Merge(br)

	bl := h.LeftAtrium.Blood.Extract(0.95)
	h.leftVentricle.Blood.Merge(bl)

	// ventricles pump out blood to their outtakes
	bro := h.rightVentricle.Blood.Extract(0.95)
	h.rightVentricle.Artery.AcceptBlood(bro)

	blo := h.leftVentricle.Blood.Extract(0.95)
	h.leftVentricle.Artery.AcceptBlood(blo)

	// set new target heart rate based on myocardial blood vessel's neurotransmitters.
	return h.pacemaker.targetHeartRate(h.Myocardium.Norepinephrine())
}

func (p *pacemaker) targetHeartRate(norepinephrine float64) float64 {
	rest := 50.0
	return rest + math.Floor(norepinephrine*130)
}

func (h *Heart) MonitorHeart() []*simulation.BloodStatistics {
	stats := []*simulation.BloodStatistics{}

	totalQty := h.LeftAtrium.Blood.Quantity +
		h.RightAtrium.Blood.Quantity +
		h.leftVentricle.Blood.Quantity +
		h.rightVentricle.Blood.Quantity +
		h.Myocardium.BloodQuantity()

	stats = append(stats, &simulation.BloodStatistics{
		ComponentName: "heart",
		BloodQuantity: totalQty,
	})

	return stats
}
