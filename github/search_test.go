//https://github.com/xiaolai
//https://github.com/ruanyf
//https://github.com/ouyangzhiping

package github

import (
	"log"
	"testing"
	"fmt"
)

func TestSearch(t *testing.T) {
	var args = [][]string{
		{
			"test1",
			"repo:xibude/gotogo",
			"is:open",
			"json",
			"decoder",
		},
		{
			"test2",
			"repo:golang/go",
			"is:open",
			"json",
			"decoder",
		},
	}

	for _, v := range args {
		pResult, err := SearchIssues(v[1:])
		if err != nil {
			log.Fatal(err)
		}
		printIssues(pResult)
	}

}

func printIssues(pResult *IssuesSearchResult) {
	fmt.Printf("\n#[%d issues]:\n", pResult.TotalCount)
	fmt.Printf("#Num   | Login    | Title \n")
	for _, item := range pResult.Items {
		fmt.Printf("#%-5d | %9.9s | %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
