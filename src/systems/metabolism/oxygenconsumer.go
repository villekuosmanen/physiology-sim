package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

type OxygenConsumer struct{}

var _ Metaboliser = (*OxygenConsumer)(nil)

const oxygenConsumptionRate = 0.000101

// Metabolise implements Metaboliser.
func (c *OxygenConsumer) Metabolise(b *circulation.Blood) {
	current := b.OxygenSaturation

	efficiency := 1.0
	if current < 0.5 {
		// no metabolism
		return
	} else if current < 0.7 {
		efficiency = 0.1
	} else if current < 0.75 {
		efficiency = 0.2
	} else if current < 0.8 {
		efficiency = 0.35
	} else if current < 0.85 {
		efficiency = 0.5
	} else if current < 0.90 {
		efficiency = 0.8
	}

	b.OxygenSaturation = current - (oxygenConsumptionRate * current * float64(efficiency))
}
