// Issues prints a table of GitHub issues matching the search terms
package main

import (
	"fmt"
	"go-learning/src/go-programming-language/chapter4/github/githubp"
	"log"
	"os"
)

func main() {
	result, err := githubp.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
