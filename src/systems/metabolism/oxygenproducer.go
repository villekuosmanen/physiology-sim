package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type LungMetaboliser struct {
	RespitoryRate float64 // represents how fast oxygen saturation recoverss
}

var _ Metaboliser = (*LungMetaboliser)(nil)

func NewLungMetaboliser() *LungMetaboliser {
	return &LungMetaboliser{
		RespitoryRate: 1.5,
	}
}

// Metabolise implements Metaboliser.
func (p *LungMetaboliser) Metabolise(b *circulation.Blood) {
	adjustedRespitoryRate := p.RespitoryRate * b.Norepinephrine
	if adjustedRespitoryRate > 1.5 {
		adjustedRespitoryRate = 1.5
	} else if adjustedRespitoryRate < 0.1 {
		adjustedRespitoryRate = 0.1
	}
	b.OxygenSaturation = b.OxygenSaturation + (1-b.OxygenSaturation)*adjustedRespitoryRate
}
