package main

import (
	"context"
	"os"
	"time"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(ctx, os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)

	// Create a competiton
	competition, _ := repo.InsertCompetition(ctx, repository.InsertCompetitionParams{
		Name:      "SomeComp",
		StartTime: time.Now(),
		MinRounds: 3,
	})

	// Create two users
	bob, _ := repo.InsertUser(ctx, repository.InsertUserParams{
		Name:  "Bob",
		Email: "bob@example.com",
		Crypt: "123",
	})
	alex, _ := repo.InsertUser(ctx, repository.InsertUserParams{
		Name:  "alex",
		Email: "alex@example.com",
		Crypt: "1234",
	})

	// Create two inscriptions
	repo.InsertInscription(ctx, repository.InsertInscriptionParams{
		Competition: competition.ID,
		Participant: bob.ID,
	})
	repo.InsertInscription(ctx, repository.InsertInscriptionParams{
		Competition: competition.ID,
		Participant: alex.ID,
	})

	// Create a task
	task, err := repo.InsertTask(ctx, repository.InsertTaskParams{
		Name:    "Some Task",
		Details: "These are some details",
		Points:  3,
		Duration: pgtype.Interval{
			Microseconds: 1000,
			Valid:        true,
		},
	})
	if err != nil {
		panic(err)
	}

	// Create a task order
	repo.InsertTaskOrder(ctx, repository.InsertTaskOrderParams{
		Competition: competition.ID,
		Task:        task.ID,
	})
}
