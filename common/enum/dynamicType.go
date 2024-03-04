package enum

type DynamicType int

const (
	Hot DynamicType = iota
	Latest
	Recent
)
