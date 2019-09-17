package mtrf

type Mode uint8
type Ctr uint8
type Device uint8

const (
	ModeTX      Mode = 0
	ModeRX           = 1
	ModeTXF          = 2
	ModeRXF          = 3
	ModeService      = 4
	ModeUpgrade      = 5
)

const (
	CtrSend         Ctr = 0
	CtrBroadcast        = 1
	CtrRecv             = 2
	CtrBindOn           = 3
	CtrBindOff          = 4
	CtrClearChannel     = 5
	CtrClearAll         = 6
	CtrUnBind           = 7
	CtrSendF            = 8
)

const (
	DeviceUnknown Device = 0
	DevicePT111          = 1
	DevicePT112          = 2
)
