package organ

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type Heart struct {
	// contains a muscle
	Myocardium Muscle

	// contains two atria
	// contains two ventricles
	RightAtrium    atrium    // from systemic
	LeftAtrium     atrium    // from pulmonary
	rightVentricle ventricle // to pulmonary
	leftVentricle  ventricle // to systemic
}

func ConstructHeart() *Heart {
	leftAtrium := atrium{}
	rightAtrium := atrium{}

	myocardium := ConstructMuscle(&rightAtrium)

	leftVentricle := ventricle{
		// artery to be set below
		Blood: &circulation.Blood{
			Quantity: 50,
		},
	}
	rightVentricle := ventricle{
		// artery to be set below
		Blood: &circulation.Blood{
			Quantity: 50,
		},
	}

	return &Heart{
		Myocardium:     *myocardium,
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
	Blood circulation.Blood
}

var _ circulation.BloodConsumer = (*atrium)(nil)

func (a *atrium) AcceptBlood(b circulation.Blood) {
	a.Blood.Merge(b)
}

type ventricle struct {
	Blood  *circulation.Blood
	Artery circulation.BloodConsumer
}

func (h *Heart) Beat() {
	// TODO
	// blood moves from atria to ventricles
	br := h.RightAtrium.Blood.Extract(0.95)
	h.rightVentricle.Blood.Merge(br)

	bl := h.LeftAtrium.Blood.Extract(0.95)
	h.rightVentricle.Blood.Merge(bl)

	// ventricles pump out blood to their outtakes
	bro := h.rightVentricle.Blood.Extract(0.95)
	h.rightVentricle.Artery.AcceptBlood(bro)

	blo := h.rightVentricle.Blood.Extract(0.95)
	h.leftVentricle.Artery.AcceptBlood(blo)
}
