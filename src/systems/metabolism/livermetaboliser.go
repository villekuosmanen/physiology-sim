package metabolism

import (
	"math"

	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
)

// LiverMetaboliser is a metaboliser for the liver.
// It is only capable of metabolism under aerobic conditions.
// It can also metabolise lactic acid.
type LiverMetaboliser struct{}

var _ Metaboliser = (*LiverMetaboliser)(nil)

// Metabolise implements Metaboliser.
func (c *LiverMetaboliser) Metabolise(b *circulation.Blood) {
	current := b.OxygenSaturation

	powerDemand := (oxygenConsumptionRate) * 0.94 // acceptable scale factor
	aerobicProduction := oxygenConsumptionRate * current * current
	if aerobicProduction >= powerDemand {
		// use what was required only
		b.OxygenSaturation = current - powerDemand

		// burn lactic acid - liver is more efficient than muscles
		excess := aerobicProduction - powerDemand
		b.LacticAcid -= (excess * lacticAcidBurnRate * 4)
		if b.LacticAcid < 0 {
			// ensure it doesn't go below zero
			b.LacticAcid = 0
		}
		return
	}

	// just use what you can produce
	if !math.IsNaN(aerobicProduction) {
		b.OxygenSaturation = current - aerobicProduction
	}
}
