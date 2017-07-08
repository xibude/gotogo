package github

import (
	"os"
	"log"
	"testing"
)

func TestIssuesHtml(t *testing.T) {
	var args = [][]string{
		{
			"test2",
			"repo:golang/go",
			"commenter:gopherbot",
			"json",
			"decoder",
		},
	}

	for _, v := range args {
		pResult, err := SearchIssues(v[1:])
		if err != nil {
			lg.Printf("search fail %s", err)
			log.Fatal(err)
		}

		if err := issuesList.Execute(os.Stdout, pResult); err != nil {
			lg.Printf("execute fail %s", err)
			log.Fatal(err)
		}
	}
}
