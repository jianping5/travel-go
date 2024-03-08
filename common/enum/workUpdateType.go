package enum

type WorkUpdateType int

const (
	Remove WorkUpdateType = iota
	Sell
	Buy
)
