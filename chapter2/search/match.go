package search

import (
	"fmt"
	//"github.com/astaxie/beego/logs"
	"log"
	"os"
)

const resultFile = "data/result"

type Result struct {
	Title   string
	Content string
}

func (r *Result) String() string {
	return fmt.Sprintf("%s :%s \n\n", r.Title, r.Content)

}

type Matcher interface {
	Search(feed *Feed, searchKey string) ([]*Result, error)
}

func Match(feed *Feed, searchKey string, matcher Matcher, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchKey)

	if err != nil {
		log.Println(err.Error())
		return

	}

	for _, result := range searchResults {
		results <- result
	}
}

func SaveResult(results chan *Result) {

	file, err := os.OpenFile(resultFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		log.Fatal("Open result file failed:", err)

	}

	defer file.Close()

	for result := range results {

		n, err := file.WriteString(result.String())
		log.Println(n, err)
	}
}

func Dispaly(results chan *Result) {
	for result := range results {
		log.Println(result.String())
	}
}
