package crons

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/internal/matchmaking"
	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
	"github.com/sirupsen/logrus"
)

const matchMakingDeltaTime = time.Second * 3

func BackgroundMatchMaking(ctx context.Context, repo *repository.Queries) {
	ticker := time.NewTicker(matchMakingDeltaTime)
	defer ticker.Stop()

	for range ticker.C {
		err := BackgroundMatchMakingStep(ctx, repo)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func BackgroundMatchMakingStep(ctx context.Context, repo *repository.Queries) error {
	err := StartPendingCompetitions(ctx, repo)
	if err != nil {
		return err
	}
	err = ProcessCurrentlyRunningCompetitons(ctx, repo)
	if err != nil {
		return err
	}

	return nil
}

func StartPendingCompetitions(ctx context.Context, repo *repository.Queries) error {
	// Check for pending competitions and start them
	competitions, err := repo.FindAllCompetitionsToStart(ctx)
	if err != nil {
		return fmt.Errorf("Failed to find comepetitions to start due to: %w", err)
	}
	for _, competition := range competitions {
		log.Printf("Starting competition %v", competition.Name)
		_, err := repo.SetCompetitionStatus(ctx, repository.SetCompetitionStatusParams{
			ID:     competition.ID,
			Status: repository.UnitStatusRunning,
		})
		if err != nil {
			logrus.Errorln(fmt.Errorf("Failed to start competition %s due to: %w", competition.ID, err))
			continue
		}

		_, err = matchmaking.GenerateInitialMatches(ctx, repo, competition.ID)
		if err != nil {
			return fmt.Errorf("Failed to generate initial matches for %s due to: %w", competition.ID, err)
		}
	}

	// TODO: Check the start channel

	return nil
}
func ProcessCurrentlyRunningCompetitons(ctx context.Context, repo *repository.Queries) error {
	// Get all currently running competitions
	competitions, err := repo.FindAllRunningCompetitions(ctx)
	if err != nil {
		return err
	}
	log.Printf("The amount of currently running competitions is: %v", len(competitions))
	for _, competition := range competitions {
		// Process running matches
		runningMatches, err := repo.FindAllRunningMatchesInCompetition(ctx, repository.FindAllRunningMatchesInCompetitionParams{
			Competition: competition.ID,
		})
		err = ProcessRunningMatchesOfCompetition(ctx, repo, competition, runningMatches)
		if err != nil {
			return err
		}
	}

	// Process completed matches
	err = ProcessCompletedMatchesOfCompetitions(ctx, repo)
	if err != nil {
		return err
	}

	return err
}

func ProcessCompletedMatchesOfCompetitions(ctx context.Context, repo *repository.Queries) error {
	stats, err := repo.GetCompetitionDescendentlessMatchStats(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get stats for descendentless matches due to: %w", err)
	}

	for _, stat := range stats {
		log.Println("=======================")
		log.Printf("Stats of %v:", stat.Competition)
		log.Printf("CompletedCount: %v", stat.CompletedCount)
		log.Printf("TotalCount: %v", stat.TotalCount)
		log.Println("=======================")

		if stat.CompletedCount == stat.TotalCount {
			if stat.TotalCount == 1 {
				err := repo.FinishCompetition(ctx, repository.FinishCompetitionParams{
					ID: stat.Competition,
				})
				if err != nil {
					return fmt.Errorf("Failed to finish competition %v due to: %w", stat.Competition, err)
				}
			} else {
				_, err := matchmaking.GenerateMatchesFromFinishedOnes(ctx, repo, stat.Competition)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func ProcessRunningMatchesOfCompetition(ctx context.Context, repo *repository.Queries, competition repository.Competition, matches []repository.Match) error {
	// Find All Running Matches
	for _, match := range matches {
		if match.Status == repository.UnitStatusRunning {
			// Find All Compeleted Rounds
			rounds, err := repo.FindAllCompletedRoundsInMatch(ctx, repository.FindAllCompletedRoundsInMatchParams{
				Match: match.ID,
			})
			if err != nil {
				return fmt.Errorf("Failed to find completed rounds for match %v due to: %w", match.ID, err)
			}

			// Check if min_rounds is satisfied
			if len(rounds) < int(competition.MinRounds) {
				// If not, check if there are any currently running matches
			} else {
				// Check if a tie break is required

			}

		}
	}

	// Spawn a new round if needed

	// Or... Set the winner and finish the match

	return nil
}
