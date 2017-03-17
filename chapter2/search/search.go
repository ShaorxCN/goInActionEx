package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func RegisterMatcher(feedtype string, matcher Matcher) {
	if _, exists := matchers[feedtype]; exists {
		log.Fatalf("%s exists in matchers", feedtype)

	}

	matchers[feedtype] = matcher
	log.Println(feedtype, "Matcher registered")
}

func Run(searchKey string) {

	feeds, err := GetDataSource()

	if err != nil {
		log.Panicln("Get source failed:", err)
	}

	results := make(chan *Result)

	var wg sync.WaitGroup

	wg.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {

			Match(feed, searchKey, matcher, results)
			wg.Done()
		}(matcher, feed)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	//Dispaly(results)
	SaveResult(results)
}
