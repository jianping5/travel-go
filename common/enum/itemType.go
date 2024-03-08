package enum

type ItemType int

const (
	VIDEO ItemType = iota
	ARTICLE
	DYNAMIC
	COMMENT
	USER
	COPYRIGHT
	COMMUNITY
)
