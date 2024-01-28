package simulation

import "fmt"

type BloodStatistics struct {
	ComponentName       string  `json:"component_name"`
	BloodQuantity       float64 `json:"blood_quantity"`
	HasOxygenSaturation bool    `json:"has_oxygen_saturation"`
	OxygenSaturation    float64 `json:"oxygen_saturation"`
	HasLacticAcid       bool    `json:"has_lactic_acid"`
	LacticAcid          float64 `json:"lactic_acid"`
	HasNorepinephrine   bool    `json:"has_norepinephrine"`
	Norepinephrine      float64 `json:"norepinephrine"`
	Verbose             bool    `json:"verbose"`
}

type GeneralStatistics struct {
	HeartRate float64 `json:"heart_rate"`
	Effort    float64 `json:"effort"`
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
