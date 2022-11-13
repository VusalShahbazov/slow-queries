package service

import (
	"context"

	"github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries"
)

type SlowQueriesService struct {
	repo slow_queries.Repository
}

func (s SlowQueriesService) GetQueries(ctx context.Context, filter slow_queries.QueryFilter) (slow_queries.QueryResult, error) {
	return s.repo.GetQueries(ctx, filter)
}

func New(repo slow_queries.Repository) slow_queries.Service {
	return &SlowQueriesService{repo}
}
