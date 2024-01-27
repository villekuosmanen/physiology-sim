package circulation

type BloodConsumer interface {
	AcceptBlood(Blood)
}

type Vessel struct {
	// TODO
	// - contains a reservoir for blood
	// - at each tick, a share of the blood avaiable in the artery is sent to its outflows
	Blood Blood
}
