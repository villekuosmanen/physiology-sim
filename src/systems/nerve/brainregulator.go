package nerve

import "github.com/villekuosmanen/physiology-sim/src/systems/circulation"

// BrainRegulator releases or neutralises neurotransmitters based on blood readings.
type BrainRegulator struct {
	previousLacticAcid       float64
	previousOxygenSaturation float64
}

const norepinephrineReleaseRate = 0.2
const norepinephrineDeactivationRate = 0.05

func (r *BrainRegulator) Regulate(bl circulation.Blood) circulation.Blood {
	diff := r.norepinephrineChange(bl)
	if diff > 0 {
		bl.Norepinephrine += norepinephrineReleaseRate
	} else {
		bl.Norepinephrine -= norepinephrineDeactivationRate
	}

	if bl.Norepinephrine < 0 {
		// can not go negative
		bl.Norepinephrine = 0
	} else if bl.Norepinephrine > 1 {
		bl.Norepinephrine = 1
	}

	r.previousLacticAcid = bl.LacticAcid
	r.previousOxygenSaturation = bl.OxygenSaturation

	return bl
}

func (r *BrainRegulator) norepinephrineChange(bl circulation.Blood) float64 {
	if (bl.LacticAcid + r.previousLacticAcid) == 0 {
		// no lactic acid - remove Norepinephrine
		return 0
	}
	lacticDiff := (bl.LacticAcid - r.previousLacticAcid) / (bl.LacticAcid + r.previousLacticAcid)
	oxyDiff := ((1 - bl.OxygenSaturation) - (1 - r.previousOxygenSaturation)) / ((1 - bl.OxygenSaturation) + (1 - r.previousOxygenSaturation))

	if bl.LacticAcid < 0.05 && bl.OxygenSaturation > 0.92 {
		return 0
	}
	return lacticDiff + oxyDiff
}
