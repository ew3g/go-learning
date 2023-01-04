// Issues prints a table of GitHub issues matching the search terms
package main

import (
	"go-learning/src/go-programming-language/chapter4/github/githubp"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues: 
	{{range .Items}}---------------------------------------------------
	Number: {{.Number}}
	User:   {{.User.Login}}
	Title:  {{.Title | print "%.64s"}}
	Age:    {{.CreatedAt | daysAgo}} days
	{{end}}`

var report, err = template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ)

func main() {
	if err != nil {
		log.Fatal(err)
	}

	result, err := githubp.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
