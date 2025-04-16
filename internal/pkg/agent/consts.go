package agent

// mode
const (
	BasicMode  = "basic"
	RemoteMode = "remote"
	LocalMode  = "local"
)

// pool param
const (
	MinLen = 1
	MaxLen = 10

	Seq          = 1
	Parallel     = 2
	Mixed        = 3
	ReverseMixed = 4
)
