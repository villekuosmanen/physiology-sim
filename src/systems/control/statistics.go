package control

import "fmt"

type BloodStatistics struct {
	ComponentName       string
	BloodQuantity       float64
	HasOxygenSaturation bool
	OxygenSaturation    float64
}

// Print implements StatisticCarrier
func (s *BloodStatistics) Print() {
	str := fmt.Sprintf("%s:\n  - Blood count: %.2f\n", s.ComponentName, s.BloodQuantity)
	if s.HasOxygenSaturation {
		str += fmt.Sprintf("  - Oxygen saturation: %.2f\n", s.OxygenSaturation)
	}

	fmt.Print(str)
}
