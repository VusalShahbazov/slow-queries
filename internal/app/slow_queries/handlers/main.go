package handlers

import (
	"context"

	"github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries"
	"github.com/VusalShahbazov/slow-query-logs/internal/domains"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SlowQueriesHandler struct {
	srv       slow_queries.Service
	validator *validator.Validate
}

func Init(app *fiber.App, srv slow_queries.Service) {
	h := SlowQueriesHandler{
		srv:       srv,
		validator: validator.New(),
	}

	app.Get("/slow_queries", h.GetQueries)
}

type GetQueriesRequest struct {
	Type       string `query:"type"`
	SortColumn string `query:"sort_column" validate:"omitempty,oneof=mean_time max_time"`
	SortType   string `query:"sort_type" validate:"omitempty,oneof=desc asc"`
	Page       int    `query:"page" validate:"omitempty,gte=0"`
	PerPage    int    `query:"per_page" validate:"omitempty,gte=0"`
}

func (q *GetQueriesRequest) page() int {
	if q.Page <= 0 {
		return 0
	}
	return q.Page - 1
}

func (q *GetQueriesRequest) perPage() int {
	if q.PerPage <= 0 {
		return 10
	}

	return q.PerPage
}

type GetQueriesResponse struct {
	Page  int                 `json:"page"`
	Total int                 `json:"total"`
	Items []domains.SlowQuery `json:"items"`
}

func (s *SlowQueriesHandler) GetQueries(ctx *fiber.Ctx) error {
	req := GetQueriesRequest{}

	// request bind
	err := ctx.QueryParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(map[string]string{"error": "invalid input"})
	}

	if err := s.validator.Struct(&req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return ctx.Status(400).JSON(map[string]string{"error": validationErrors.Error()})
	}

	filter := slow_queries.QueryFilter{
		Type:       req.Type,
		SortColumn: req.SortColumn,
		SortType:   req.SortType,
		Limit:      req.perPage(),
		Offset:     req.page() * req.perPage(),
	}

	result, err := s.srv.GetQueries(context.Background(), filter)
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"error": err.Error()})
	}

	res := GetQueriesResponse{
		Page:  req.page() + 1,
		Total: result.Count,
		Items: result.Items,
	}

	return ctx.Status(200).JSON(res)
}
