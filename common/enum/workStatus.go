package enum

type WorkStatus int

const (
	Created WorkStatus = iota
	OnSale
	Sold
)
