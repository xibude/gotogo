// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package douban

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	//"strings"
	"log"
	"os"
)

var Lg = log.New(os.Stdout, "<whl>", log.Lshortfile | log.LstdFlags)

// SearchIssues queries the GitHub issue tracker.
func SearchBook(bookId string) (*string, error) {

	Lg.Println("url is:", BookV2URL)
	resp, err := http.Get(BookV2URL)
	if err != nil {
		return nil, err
	}
	Lg.Println("resp is :", resp)
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	Lg.Println("body is:", resp.Body)

	var result []*bookInfo
	dec := json.NewDecoder(resp.Body)
	Lg.Println("dec is:", dec)

	if err := dec.Decode(&result); err != nil {
		Lg.Printf("decode err=", err)
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	Lg.Println("result is:",result)

	for _,v := range result {
		Lg.Printf("---%+v", v)
	}
	//return &result, nil
	return nil, nil
}

//!-