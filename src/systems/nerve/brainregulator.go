package nerve

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

// BrainRegulator releases or neutralises neurotransmitters based on blood readings.
type BrainRegulator struct{}

const norepinephrineReleaseRate = 0.01
const norepinephrineDeactivationRate = 0.001

func (r *BrainRegulator) Regulate(bl circulation.Blood) circulation.Blood {
	acidity := bl.Acidity()
	if acidity > 0 {
		bl.Norepinephrine += acidity * norepinephrineReleaseRate
	} else {
		bl.Norepinephrine += acidity * norepinephrineDeactivationRate
	}

	return bl
}
