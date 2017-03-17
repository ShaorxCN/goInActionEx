package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	RegisterMatcher("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchKey string) ([]*Result, error) {
	return nil, nil
}
