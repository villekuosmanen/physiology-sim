package organ

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type Heart struct {
	// contains a muscle
	Myocardium Muscle

	// contains two atria
	// contains two ventricles
	rightAtrium    atrium    // from systemic
	rightVentricle ventricle // to pulmonary
	leftAtrium     atrium    // from pulmonary
	leftVentricle  ventricle // to pulmonory
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
	br := h.rightAtrium.Blood.Extract(0.95)
	h.rightVentricle.Blood.Merge(br)

	bl := h.leftAtrium.Blood.Extract(0.95)
	h.rightVentricle.Blood.Merge(bl)

	// ventricles pump out blood to their outtakes
	bro := h.rightVentricle.Blood.Extract(0.95)
	h.rightVentricle.Artery.AcceptBlood(bro)

	blo := h.rightVentricle.Blood.Extract(0.95)
	h.leftVentricle.Artery.AcceptBlood(blo)
}
