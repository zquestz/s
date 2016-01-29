package arxiv

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
		// Single word.
		{"networks", fmt.Sprintf(
			"http://arxiv.org/find/all/1/all:+%s/0/1/0/all/0/1", "networks")},
		// Two words.
		{"convolutional networks", fmt.Sprintf(
			"http://arxiv.org/find/all/1/all:+%s/0/1/0/all/0/1",
			"AND+convolutional+networks")},
		// More than two words.
		{"deep convolutional networks", fmt.Sprintf(
			"http://arxiv.org/find/all/1/all:+%s/0/1/0/all/0/1",
			"AND+networks+AND+deep+convolutional")},
		// Query with or.
		{"machine learning or deep convolutional networks", fmt.Sprintf(
			"http://arxiv.org/find/all/1/all:+%s/0/1/0/all/0/1",
			"OR+AND+machine+learning+AND+networks+AND+deep+convolutional")},
		// Query with multiple or.
		{"machine learning or deep neural networks or language processing",
			fmt.Sprintf("http://arxiv.org/find/all/1/all:+%s/0/1/0/all/0/1",
				"OR+AND+language+processing+OR+AND+machine+learning+"+
					"AND+networks+AND+deep+neural")},
	}
	for _, c := range cases {
		got := p.BuildURI(c.query)
		if got != c.want {
			t.Errorf("BuildURI(%q) == %q, want %q", c.query, got, c.want)
		}
	}
}
