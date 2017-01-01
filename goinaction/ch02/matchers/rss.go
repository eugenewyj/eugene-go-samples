package matchers

import (
	"github.com/eugenewyj/go-sample/goinaction/ch02/search"
)

type rssMatcher struct {
}

func init()  {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (rss rssMatcher) Search (feed *search.Feed, searchTerm string) ([]*search.Result, error) {
 	return nil, nil
}