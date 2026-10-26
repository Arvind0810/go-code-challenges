package crawler

import "context"

// Fetcher retrieves a page and the URLs it links to.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Result records the outcome of fetching one URL. Depth is the length of the
// shortest link path from the start URL (the start URL has depth 0).
type Result struct {
	Body  string
	URLs  []string
	Depth int
	Err   error
}

// Crawl fetches startURL and every page reachable from it within maxDepth
// links, fetching each URL at most once and running at most maxConcurrency
// fetches at a time (values below 1 are treated as 1). It returns a Result
// for every URL it fetched. Once ctx is done no new fetches are started, and
// Crawl returns after in-flight fetches complete.
func Crawl(ctx context.Context, startURL string, maxDepth int, maxConcurrency int, f Fetcher) map[string]Result {
	// TODO: implement
	return nil
}
