package slow_queries

import "github.com/VusalShahbazov/slow-query-logs/internal/domains"

type QueryFilter struct {
	Type       string `json:"type"`
	SortColumn string `json:"sort_column"`
	SortType   string `json:"sort_type"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

type QueryResult struct {
	Count int
	Items []domains.SlowQuery
}
