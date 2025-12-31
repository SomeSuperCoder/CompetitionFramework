package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/app/services"
	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
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

	conn, err := pgx.Connect(ctx, os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)

	competitionService := &services.CompetitionService{Repo: repo}
	s.RegisterService(competitionService, "Competition")

	taskService := &services.TaskService{Repo: repo}
	s.RegisterService(taskService, "Task")

	matchService := &services.MatchService{Repo: repo}
	s.RegisterService(matchService, "Match")

	http.Handle("/rpc", s)

	log.Printf("RPC started and is listening on :%v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}
