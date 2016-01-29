package arxiv

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("arxiv", &ArxivProvider{})
}

// ArxivProvider adheres to the Provider interface.
type ArxivProvider struct {
}

// If the query contains more than 1 word, the format to binary logical
// combination as follows:
// "neural networks" -> "OP neural networks"
// "convolutional neural networks" -> "OP networks OP convolutional neural"
func formatWithOp(qs []string, op string) string {
	var formatted_string string

	if len(qs) == 1 {
		formatted_string = qs[0]
	} else {
		formatted_string = fmt.Sprintf("%s %s", op, strings.Join(qs[:2], " "))
	}
	// Populate all the additional entries with OP in front of the current
	// formatted string.
	for i := 2; i < len(qs); i++ {
		formatted_string = fmt.Sprintf("%s %s %s", op, qs[i], formatted_string)
	}
	return formatted_string
}

// BuildURI generates a search URL for ArXiv.
func (p *ArxivProvider) BuildURI(q string) string {
	// Separate query by "or".
	queries := strings.Split(q, " or ")
	results := make([]string, len(queries))
	for i := 0; i < len(queries); i++ {
		// Each query should be formatted with "AND".
		splittedQuery := strings.Split(queries[i], " ")
		results[i] = formatWithOp(splittedQuery, "AND")
	}

	return fmt.Sprintf("http://arxiv.org/find/all/1/all:+%s/0/1/0/all/0/1",
		url.QueryEscape(formatWithOp(results, "OR")))
}
