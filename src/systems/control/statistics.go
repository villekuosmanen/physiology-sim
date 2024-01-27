package control

type Statistics interface {
	Print()
}

type BloodStatistics struct {
	BloodQuantity float64
}
