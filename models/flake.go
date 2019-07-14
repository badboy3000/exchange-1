package models

// Flake {"id":257532141091473317,"machine-id":51109,"msb":0,"sequence":0,"time":15350111788}
type Flake struct {
	ID        uint64 `json:"id"`
	MachineID uint64 `json:"machine-id"`
	Msb       uint64 `json:"msb"`
	Sequence  uint64 `json:"sequence"`
	Time      uint64 `json:"time"`
}
