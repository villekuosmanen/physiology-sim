package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

// OxygenConsumer is a simple oxygen consumer.
// It is only capable of metabolism under aerobic conditions.
type OxygenConsumer struct{}

var _ Metaboliser = (*OxygenConsumer)(nil)

const oxygenConsumptionRate = 0.0001

// Metabolise implements Metaboliser.
func (c *OxygenConsumer) Metabolise(b *circulation.Blood) {
	current := b.OxygenSaturation

	powerDemand := (oxygenConsumptionRate) * 0.92 // acceptable scale factor
	aerobicProduction := oxygenConsumptionRate * current * current
	if aerobicProduction >= powerDemand {
		// use what was required only
		b.OxygenSaturation = current - powerDemand
		return
	}

	// just use what you can produce
	b.OxygenSaturation = current - aerobicProduction
}
