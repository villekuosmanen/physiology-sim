package metabolism

type MET float64 // Metabolic Equivalent of Task

const (
	METRest          MET = 0.66 // METRest decribes metabolic rate at rest
	METSitting       MET = 1
	METLightCardio   MET = 3
	METMediumCardio  MET = 5.5
	METHardCardio    MET = 8
	METExtremeCardio MET = 10
)

func (m MET) Float64() float64 {
	return float64(m)
}
