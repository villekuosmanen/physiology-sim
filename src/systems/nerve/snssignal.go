package nerve

type SNSSignalHandleMethod int

const (
	SNSSignalHandleMethodNothing = iota + 1
	SNSSignalHandleMethodExpand
	SNSSignalHandleMethodContract
)
