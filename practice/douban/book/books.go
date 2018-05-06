// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
// DouBan: https://developers.douban.com/wiki/?title=api_v2
package book

const BookV2URL = "https://api.douban.com/v2/book/"

type BookInfo struct {
	Rating *Rating
	Tags []*Tags
	Pages string
	ID string
	Title string
	Price string
}

type Rating struct {
	Max uint32
	NumRaters uint32
	Average string
	Min uint32
}

type Tags struct {
	Count uint32
	Name string
	Title string
}




