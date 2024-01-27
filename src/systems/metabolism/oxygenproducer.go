package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type OxygenProducer struct {
	RateFactor float64 // represents how fast oxygen saturation recovers
}

var _ Metaboliser = (*OxygenProducer)(nil)

// Metabolise implements Metaboliser.
func (p *OxygenProducer) Metabolise(b *circulation.Blood) {
	b.OxygenSaturation = b.OxygenSaturation + (1-b.OxygenSaturation)*p.RateFactor
}
