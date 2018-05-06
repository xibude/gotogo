//https://github.com/xiaolai
//https://github.com/ruanyf
//https://github.com/ouyangzhiping

package github

import (
	"log"
	"testing"
	"fmt"
	"time"
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
		printIssuesSort(pResult)
	}

}

func printIssues(pResult *IssuesSearchResult) {
	fmt.Printf("\n#[%d issues]: now is %+v\n",
		pResult.TotalCount, time.Now().Format("2006-01-03 15:04:05"))
	fmt.Printf("#Num   | Login    | Title \n")
	for _, item := range pResult.Items {

			fmt.Printf("#%-5d | %9.9s | %+v | %.55s\n",
				item.Number, item.User.Login, item.CreateAt, item.Title)

	}
	fmt.Printf("\n")

}

var timeBegin = "2017-01-01 00:00:00"

func printIssuesSort(pResult *IssuesSearchResult) {
	tBegin, _ := time.Parse("2006-01-02 15:04:05", timeBegin)

	fmt.Printf("\n#[%d issues]: now is %+v\n",
		pResult.TotalCount, time.Now().Format("2006-01-03 15:04:05"))
	fmt.Printf("#Num   | Login    | Title \n")
	for _, item := range pResult.Items {
		if item.CreateAt.Before(tBegin) {
			fmt.Printf("#%-5d | %9.9s | %+v | %.55s\n",
				item.Number, item.User.Login, item.CreateAt, item.Title)
		}
	}
	fmt.Printf("\n")
	for _, item := range pResult.Items {
		if item.CreateAt.After(tBegin) {
			fmt.Printf("#%-5d | %9.9s | %+v | %.55s\n",
				item.Number, item.User.Login, item.CreateAt, item.Title)
		}
	}
}


