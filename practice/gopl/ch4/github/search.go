// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"errors"
	"log"
	"os"
	//"gopl.io/ch4/github"
)

var lg *log.Logger = log.New(os.Stdout,"", log.Lshortfile | log.LstdFlags)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	aa := 1
	var resp *http.Response
	var err = errors.New("")

	lg.Printf("terms=%v", terms)
	q := url.QueryEscape(strings.Join(terms, " "))

	lg.Printf("q=%v", q)
	//./issues repo:golang/go is:open json decoder

	lg.Printf("URL=%v", IssuesURL+"?q="+q)
	//https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
	if aa == 0 {
		resp, err = http.Get(IssuesURL + "?q=" + q)
		if err != nil {
			return nil, err
		}
	}

	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	if aa == 1 {
		req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set(
			"Accept", "application/vnd.github.v3.text-match+json")
		resp, err = http.DefaultClient.Do(req)
	}

	//lg.Printf("resp=%+v", resp)
	lg.Printf("resp.body=%+v", resp.Body)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	lg.Printf("result=%+v. \r\n", &result)

	for _, v := range result.Items {
		lg.Printf("---num=%v, color=%v", v.Number, v.Labels[0].Color)
	}

	resp.Body.Close()
	return &result, nil
}

func PrintIssues(terms []string) {
	result, err := SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)

		fmt.Printf("  Label<ID=%s, Name=%s Color=%s\n",
			item.Labels[0].ID, item.Labels[0].Name, item.Labels[0].Color)
	}
}
//!-
