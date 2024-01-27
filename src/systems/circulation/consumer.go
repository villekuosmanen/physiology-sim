package circulation

type BloodConsumer interface {
	AcceptBlood(Blood)
}
