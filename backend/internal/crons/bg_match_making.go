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
		err := ProcessMatchesOfCompeition(ctx, repo, competition)
		if err != nil {
			return err
		}
	}
	// Find All Running Matches

	// Find All Running Rounds

	// Check if min_rounds is satisfied or if a tie break is required

	// Spawn a new round if needed

	// Or... Set the winner and finish the match
	return err
}

func ProcessMatchesOfCompeition(ctx context.Context, repo *repository.Queries, competition repository.Competition) error {
	matches, err := repo.FindAllCompletedMatchesInCompetitionWithNoDescendents(ctx, repository.FindAllCompletedMatchesInCompetitionWithNoDescendentsParams{
		Competition: competition.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to get matches for competition %v due to: %w", competition.Name, err)
	}

	// Check if all matches are finished
	finishedCount := CountFinishedMatches(matches)
	matchAmount := len(matches)
	if finishedCount == matchAmount {
		// End the competition if there is only one match left
		if matchAmount == 1 {
			err := FinishCompetition(ctx, repo, competition, matches)
			if err != nil {
				return err
			}
		} else {
			// Otherwise generate a new match set
			_, err := matchmaking.GenerateMatchesFromFinishedOnes(ctx, repo, matches)
			if err != nil {
				return fmt.Errorf("Failed to generate a new match set from finished ones for %s due to: %w", competition.ID, err)
			}

		}
	}
	return nil
}

func CountFinishedMatches(matches []repository.Match) int {
	finishedCount := 0
	for _, match := range matches {
		if match.Status == repository.UnitStatusCompleted {
			finishedCount += 1
		}
	}
	return finishedCount
}

func FinishCompetition(ctx context.Context, repo *repository.Queries, competition repository.Competition, matches []repository.Match) error {
	_, err := repo.SetCompetitionWinner(ctx, repository.SetCompetitionWinnerParams{
		ID:     competition.ID,
		Winner: matches[0].Winner,
	})
	if err != nil {
		return fmt.Errorf("Failed to set competition winner %v due to: %w", competition.Name, err)
	}
	_, err = repo.SetCompetitionStatus(ctx, repository.SetCompetitionStatusParams{
		ID:     competition.ID,
		Status: repository.UnitStatusCompleted,
	})
	if err != nil {
		return fmt.Errorf("Failed to finish competition %v due to: %w", competition.Name, err)
	}
	return nil
}
