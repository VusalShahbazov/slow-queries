package repository

import (
	"context"
	"fmt"

	"github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries"
	"github.com/VusalShahbazov/slow-query-logs/internal/domains"
	"gorm.io/gorm"
)

type SlowQueriesRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) slow_queries.Repository {
	return &SlowQueriesRepository{db}
}

func (r *SlowQueriesRepository) GetQueries(ctx context.Context, filter slow_queries.QueryFilter) (slow_queries.QueryResult, error) {
	var res slow_queries.QueryResult
	var total int64

	query := r.db.WithContext(ctx).Model(&domains.SlowQuery{})
	if filter.Type != "" {
		query = query.Where("lower(query) like lower(?)", filter.Type+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return res, err
	}

	if filter.SortColumn != "" {
		sortType := "desc"
		if filter.SortType != "" {
			sortType = filter.SortType
		}

		query.Order(fmt.Sprintf("%v %v", filter.SortColumn, sortType))
	}

	var data []domains.SlowQuery
	err = query.Limit(filter.Limit).Offset(filter.Offset).Find(&data).Error
	if err != nil {
		return res, err
	}

	res.Items = data
	res.Count = int(total)

	return res, nil
}
