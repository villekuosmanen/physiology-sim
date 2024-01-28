package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type LungMetaboliser struct {
	RespitoryRate float64 // represents how fast oxygen saturation recoverss
}

var _ Metaboliser = (*LungMetaboliser)(nil)

const respitoryRateScaler = 800.0

func NewLungMetaboliser() *LungMetaboliser {
	return &LungMetaboliser{
		RespitoryRate: 1,
	}
}

// Metabolise implements Metaboliser.
func (p *LungMetaboliser) Metabolise(b *circulation.Blood) {
	adjustedRespitoryRate := 1 + (p.RespitoryRate * 14)
	if adjustedRespitoryRate > 15 {
		adjustedRespitoryRate = 15
	} else if adjustedRespitoryRate < 1 {
		adjustedRespitoryRate = 1
	}
	b.OxygenSaturation = b.OxygenSaturation + (1-b.OxygenSaturation)*(adjustedRespitoryRate/respitoryRateScaler)
}
