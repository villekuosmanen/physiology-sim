package metabolism

type MET float64 // Metabolic Equivalent of Task

const (
	METRest          MET = 0.66 // METRest decribes metabolic rate at rest
	METSitting       MET = 1
	METLightCardio   MET = 3
	METMediumCardio  MET = 5.5
	METHeavyCardio   MET = 8
	METExtremeCardio MET = 10
)

func (m MET) Float64() float64 {
	return float64(m)
}

func (m MET) String() string {
	switch m {
	case METRest:
		return "Rest"
	case METSitting:
		return "Sitting"
	case METLightCardio:
		return "Light Cardio"
	case METMediumCardio:
		return "Medium Cardio"
	case METHeavyCardio:
		return "Heavy Cardio"
	case METExtremeCardio:
		return "Extreme Cardio"
	}
	return ""
}
