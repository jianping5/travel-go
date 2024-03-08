package enum

type BehaviorType int

const (
	Like BehaviorType = iota
	View
	Favor
	Comment
)
