package slow_queries

import "context"

type Repository interface {
	GetQueries(ctx context.Context, filter QueryFilter) (QueryResult, error)
}
