package control

import "fmt"

type BloodStatistics struct {
	ComponentName string
	BloodQuantity float64
}

// Print implements StatisticCarrier
func (s *BloodStatistics) Print() {
	fmt.Printf("%s:\n  - Blood count: %.2f\n", s.ComponentName, s.BloodQuantity)
}
