package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/internal/crons"
	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
	"github.com/SomeSuperCoder/CompetitionFramework/backend/rpc/services"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const port = 8899

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	// Connect to database
	conn, err := pgx.Connect(ctx, os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)

	// Register Services
	competitionService := &services.CompetitionService{Repo: repo}
	s.RegisterService(competitionService, "Competition")

	taskService := &services.TaskService{Repo: repo}
	s.RegisterService(taskService, "Task")

	matchService := &services.MatchService{Repo: repo}
	s.RegisterService(matchService, "Match")

	userService := &services.UserService{Repo: repo}
	s.RegisterService(userService, "User")

	inscriptionService := &services.InscriptionService{Repo: repo}
	s.RegisterService(inscriptionService, "Inscription")

	// Background Match making
	go crons.BackgroundMatchMaking(ctx, repo)

	// Start the http server
	http.Handle("/rpc", corsMiddleware(s))

	log.Printf("RPC started and is listening on :%v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}

// CORS middleware function
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers for all responses
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust in production
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}
