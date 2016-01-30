package reddit

import (
	"fmt"
	"testing"
)

// TestBuildURI makes sure the BuildURI method builds valid URIs.
func TestBuildURI(t *testing.T) {
	p := Provider{}
	cases := []struct {
		query, want string
	}{
		// Regular reddit search.
		{"best startups",
			fmt.Sprintf("https://www.reddit.com/search?q=%s", "best+startups")},
		// Subreddit search.
		{"/r/cscareerquestions best startups",
			fmt.Sprintf("https://www.reddit.com%s/search?q=%s&restrict_sr=on",
				"/r/cscareerquestions", "best+startups")},
		// No query subreddit search.
		{"/r/cscareerquestions",
			fmt.Sprintf("https://www.reddit.com/search?q=%s", "%2Fr%2Fcscareerquestions")},
	}
	for _, c := range cases {
		got := p.BuildURI(c.query)
		if got != c.want {
			t.Errorf("BuildURI(%q) == %q, want %q", c.query, got, c.want)
		}
	}
}
