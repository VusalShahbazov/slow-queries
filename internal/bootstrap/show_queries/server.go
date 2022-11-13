package show_queries

import (
	"log"
	"time"

	slowQueriesHandler "github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries/handlers"
	slowQueriesRepo "github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries/repository"
	slowQueriesSrv "github.com/VusalShahbazov/slow-query-logs/internal/app/slow_queries/service"
	"github.com/VusalShahbazov/slow-query-logs/pkg/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	*Config
	validator validator.Validate
}

func (s Server) Run() error {

	app := fiber.New()

	//Just for waiting pg up can be removed
	time.Sleep(time.Second * 2)

	// Connect to postgres
	db, err := postgres.Connect(s.DBHost, s.DBPort, s.DBUser, s.DBPassword, s.DBName)
	if err != nil {
		return err
	}

	log.Printf("Connected to database successfully")

	//Init services and repos
	repo := slowQueriesRepo.New(db)
	srv := slowQueriesSrv.New(repo)

	//Init routes
	slowQueriesHandler.Init(app, srv)

	log.Printf("Start listen on:  %v \n", s.BindAddr)
	return app.Listen(s.BindAddr)
}

func NewServer(cnf *Config) *Server {
	return &Server{
		Config: cnf,
	}
}
