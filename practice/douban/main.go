// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	//"fmt"
	"log"
	//"os"

	"woohello.ipo/douban/data"
	//"strings"
)


func main() {
	_, err := douban.SearchBook("1220562")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("result is :%s\n", result)
	//fmt.Printf("%d issues:\n", result.TotalCount)

}