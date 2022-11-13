package slow_queries

import "context"

type Service interface {
	GetQueries(ctx context.Context, filter QueryFilter) (QueryResult, error)
}
