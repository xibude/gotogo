package douban

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
)

var Lg = log.New(os.Stdout, "<whl>", log.Lshortfile | log.LstdFlags)

// SearchIssues queries the douban bookinfo
func SearchBook(bookId string) (*BookInfo, error){
	if len(bookId) == 0 {
		return nil, fmt.Errorf("bookId is nil")
	}

	bookURL := BookV2URL + bookId
	Lg.Println("url is:", bookURL)

	resp, err := http.Get(bookURL)
	if err != nil {
		return nil, err
	}
	//Lg.Println("resp is :", resp)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	//Lg.Println("body is:", resp.Body)

	var result BookInfo
	dec := json.NewDecoder(resp.Body)
	//Lg.Println("dec is:", dec)

	if err := dec.Decode(&result); err != nil {
		return nil, err
	}

	//Lg.Println("result is:",result)
	return &result, nil
}

func PrintBookInfo (pBook *BookInfo) {
	fmt.Printf("\nTitle: %-32s", pBook.Title)
	fmt.Printf("\nID   : %-32s", pBook.ID)
	fmt.Printf("\nPages: %-32s", pBook.Pages)
	fmt.Printf("\nPrice: %-32s", pBook.Price)

	fmt.Printf("\n\nTagsNum: %d", len(pBook.Tags))
	for _,v := range pBook.Tags {
		fmt.Printf("\n%5d | %-32s", v.Count, v.Title)
	}
}