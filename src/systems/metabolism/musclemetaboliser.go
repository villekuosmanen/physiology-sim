package metabolism

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

// MuscleMetaboliser is a metaboliser that prefers aerobic metabolism,
// but it can also metabolise anaerobically, producing lactic acid in the process.
type MuscleMetaboliser struct {
	metabolicRate MET // MetabolicRate as measured in MET
}

const lacticAcidProductionRate = 1
const lacticAcidBurnRate = 0.5 // less efficient to burn the acid

var _ Metaboliser = (*OxygenConsumer)(nil)

func NewMuscleMetaboliser() *MuscleMetaboliser {
	return &MuscleMetaboliser{
		metabolicRate: METRest,
	}
}

func (m *MuscleMetaboliser) SetMetabolicRate(new MET) {
	m.metabolicRate = new
}

// Metabolise implements Metaboliser.
func (m *MuscleMetaboliser) Metabolise(b *circulation.Blood) {
	current := b.OxygenSaturation

	powerDemand := (oxygenConsumptionRate * m.metabolicRate.Float64()) * 0.92 // acceptable scale factor
	aerobicProduction := oxygenConsumptionRate * m.metabolicRate.Float64() * current * current
	if aerobicProduction >= powerDemand {
		// aerobic production only
		b.OxygenSaturation = current - powerDemand

		// excess power can be used to burn lactic acid
		excess := aerobicProduction - powerDemand
		b.LacticAcid -= (excess * lacticAcidBurnRate)
		if b.LacticAcid < 0 {
			// ensure it doesn't go below zero
			b.LacticAcid = 0
		}
		return
	}

	// required power will have to be produced anaerobically
	b.OxygenSaturation = current - aerobicProduction
	b.LacticAcid += (powerDemand - aerobicProduction) * lacticAcidProductionRate
}
