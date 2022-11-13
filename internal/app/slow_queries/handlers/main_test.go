package handlers_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries"
	"github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries/handlers"
	slowQueriesSrv "github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries/service"
	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestSlowQueriesHandlerGetQueries(t *testing.T) {
	testCases := map[string]TestCase{
		"success response without ant query": {
			Code:             200,
			QueryString:      "",
			ExpectedErrorKey: "",
		},
		"validation error in sort colum": {
			Code:             400,
			QueryString:      "?sort_column=wrong",
			ExpectedErrorKey: "Key: 'GetQueriesRequest.SortColumn' Error:Field validation for 'SortColumn' failed on the 'oneof' tag",
		},
		"success  in sort colum correct value ": {
			Code:             200,
			QueryString:      "?sort_column=mean_time",
			ExpectedErrorKey: "",
		},
		"validation error in sort type ": {
			Code:             400,
			QueryString:      "?sort_type=dasc",
			ExpectedErrorKey: "Key: 'GetQueriesRequest.SortType' Error:Field validation for 'SortType' failed on the 'oneof' tag",
		},
		"bind error page as a string": {
			Code:             400,
			QueryString:      "?page=string",
			ExpectedErrorKey: "invalid input",
		},
		"validation error in negative page_count": {
			Code:             400,
			QueryString:      "?per_page=-3",
			ExpectedErrorKey: "Key: 'GetQueriesRequest.PerPage' Error:Field validation for 'PerPage' failed on the 'gte' tag",
		},
	}

	app := fiber.New()

	var repo slow_queries.Repository = &mockRepo{}
	srv := slowQueriesSrv.New(repo)

	handlers.Init(app, srv)

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/slow_queries"+testCase.QueryString, nil)
			res, err := app.Test(req)
			require.NoError(t, err)

			assert.Equal(t, res.StatusCode, testCase.Code)

			if testCase.ExpectedErrorKey != "" {
				all, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				data := map[string]string{}

				err = json.Unmarshal(all, &data)
				require.NoError(t, err)

				errKey := data["error"]

				assert.Equal(t, errKey, testCase.ExpectedErrorKey)
			}
		})
	}

}

type TestCase struct {
	Code             int
	QueryString      string
	ExpectedErrorKey string
}

type mockRepo struct{}

func (m *mockRepo) GetQueries(ctx context.Context, filter slow_queries.QueryFilter) (slow_queries.QueryResult, error) {
	return slow_queries.QueryResult{}, nil
}
