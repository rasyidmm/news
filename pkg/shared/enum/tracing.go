package enum

type TracingName string

const (
	StartService	TracingName = "Start Service"
	StartInteractor TracingName = "Start Interactor"
	Error           TracingName = "Error"
	Response        TracingName = "Response"
	Request			TracingName = "Request"
)
