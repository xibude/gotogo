// Issues prints a table of GitHub issues matching the search terms.
package douban

import (
	//"fmt"
	"log"
	//"os"
	//"strings"
	"testing"
)


func TestSearchBook(t *testing.T) {
	_, err := SearchBook("1220562")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("result is :%s\n", result)
	//fmt.Printf("%d issues:\n", result.TotalCount)

}