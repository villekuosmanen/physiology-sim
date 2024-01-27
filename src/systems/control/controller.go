package control

type Controller interface {
	Act()
}

type MonitorableController interface {
	Act()
	Monitor() *BloodStatistics
}
