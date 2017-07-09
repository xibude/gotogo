package douban

import (
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	pBook, err := SearchBook("1220562")
	if err != nil {
		log.Fatal(err)
	}

	PrintBookInfo(pBook)
}
