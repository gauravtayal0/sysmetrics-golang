package constants

const (
	ActionAgentUser    = "user"
	ActionAgentMonitor = "monitor"

	JsonPrefix = ""
	JsonIndent = " "

	ChanEventMonStop = "mon-stop"

	ErrRequestRead   = "could not read request"
	ErrRequestParse  = "could not parse request"
	ErrNotFound      = "not found : %v(%v)"
	ErrAlreadyExists = "entity already exists : %v(%v)"
	ErrZookupFailed  = "zookeeper lookup failed : %v(%v)"
)
