package consts

// mode
const (
	BasicMode      = "basic"
	RemoteModeApi  = "remoteApi"
	RemoteModeFile = "remoteFile"
	LocalMode      = "local"
	MonitorMode    = "monitor"
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

const ErrKey = "err"
