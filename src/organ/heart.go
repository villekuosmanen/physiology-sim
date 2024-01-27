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

func ConstructHeart(aorta, pulmonaryArtery circulation.BloodConsumer) *Heart {
	rightAtrium := atrium{}
	leftAtrium := atrium{}

	myocardium := ConstructMuscle(&rightAtrium)

	rightVentricle := ventricle{
		Artery: pulmonaryArtery,
	}
	leftVentricle := ventricle{
		Artery: aorta,
	}

	return &Heart{
		Myocardium:     *myocardium,
		RightAtrium:    rightAtrium,
		LeftAtrium:     leftAtrium,
		rightVentricle: rightVentricle,
		leftVentricle:  leftVentricle,
	}
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
