package simulation

import "fmt"

type BloodStatistics struct {
	ComponentName       string
	BloodQuantity       float64
	HasOxygenSaturation bool
	OxygenSaturation    float64
	HasLacticAcid       bool
	LacticAcid          float64
	HasNorepinephrine   bool
	Norepinephrine      float64
	Verbose             bool
}

// Print implements StatisticCarrier
func (s *BloodStatistics) Print(verbose bool) {
	if !verbose && s.Verbose {
		// Only print verbose metrics on verbose mode
		return
	}
	str := fmt.Sprintf("%s:\n  - Blood count: %.2f\n", s.ComponentName, s.BloodQuantity)
	if s.HasOxygenSaturation {
		str += fmt.Sprintf("  - Oxygen saturation: %.2f%%\n", s.OxygenSaturation*100)
	}
	if s.HasLacticAcid {
		str += fmt.Sprintf("  - Lactic Acid: %.2f\n", s.LacticAcid*100)
	}
	if s.HasNorepinephrine {
		str += fmt.Sprintf("  - Norepinephrine: %.2f\n", s.Norepinephrine*100)
	}

	fmt.Print(str)
}
