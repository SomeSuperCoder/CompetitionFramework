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
	// ===== Competitions =========
	err := StartPendingCompetitions(ctx, repo)
	if err != nil {
		return err
	}

	// ===== Matches =========
	err = ProcessCompletedMatches(ctx, repo)
	if err != nil {
		return err
	}
	err = ProcessRunningMatches(ctx, repo)
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
func ProcessRunningMatches(ctx context.Context, repo *repository.Queries) error {
	stats, err := repo.GetLockedMatchRoundStats(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get lockend match round stats due to: %w", err)
	}
	for _, stat := range stats {
		log.Println("=======================")
		log.Printf("Stats of RUNNING MATCH %v:", stat.Match)
		log.Printf("CompletedCount: %v", stat.CompletedCount)
		log.Printf("MinRounds: %v", stat.MinRounds)
		log.Println("=======================")

		// Check if min_rounds is satisfied
		if stat.CompletedCount >= int64(stat.MinRounds) {
			// Check if a tie break is required
			if stat.User1Points == stat.User2Points {
				// TODO: spawn a new round
			} else {
				// TODO: Set the winner and finish the match
			}
		}

	}

	return nil
}

func ProcessCompletedMatches(ctx context.Context, repo *repository.Queries) error {
	stats, err := repo.GetCompetitionDescendentlessMatchStats(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get stats for descendentless matches due to: %w", err)
	}

	for _, stat := range stats {
		log.Println("=======================")
		log.Printf("Stats of RUNNING COMPETITION %v:", stat.Competition)
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
