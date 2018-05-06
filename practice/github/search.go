package github

import (
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
)

var lg *log.Logger = log.New(os.Stdout, "-->", log.Lshortfile|log.LstdFlags)

func SearchIssues(terms []string) (pResult *IssuesSearchResult, err error) {
	err = nil
	q := url.QueryEscape(strings.Join(terms, " "))
	lg.Printf("joinArgs=%s", q)

	url := IssuesURL + "?q=" + q
	lg.Printf("url=%s", url)

	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed [%d]%s", http.StatusOK, rsp.Status)
	}

	var result IssuesSearchResult
	pResult = &result

	pDecoder := json.NewDecoder(rsp.Body)
	if err = pDecoder.Decode(pResult); err != nil {
		pResult = nil
		return
	}

	return
}
