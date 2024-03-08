package enum

type ContentType int

const (
	AllContent ContentType = iota
	ForYouContent
	VideoContent
	ArticleContent
	RecentContent
)
