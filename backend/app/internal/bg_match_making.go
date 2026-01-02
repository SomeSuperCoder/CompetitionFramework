package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
	"github.com/google/uuid"
)

const matchMakingDeltaTime = time.Second * 3

func BackgroundMatchMaking(ctx context.Context, repo *repository.Queries) {
	ticker := time.NewTicker(matchMakingDeltaTime)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Running Cron Job")
	}
}

func BackgroundMatchMakingStep(ctx context.Context, repo *repository.Queries) error {
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
			return fmt.Errorf("Failed to start competition %s due to: %w", competition.ID, err)
		}

		_, err = GenerateInitialMatches(ctx, repo, competition.ID)
		if err != nil {
			return fmt.Errorf("Failed to generate initial matches for %s due to: %w", competition.ID, err)
		}
	}

	// TODO: Check the start channel

	// Get all currently running competitions
	competitions, err = repo.FindAllRunningCompetitions(ctx)
	for _, competition := range competitions {
		matches, err := repo.FindAllRunningMatchesInCompetition(ctx, repository.FindAllRunningMatchesInCompetitionParams{
			Competition: competition.ID,
		})
		if err != nil {
			return fmt.Errorf("Failed to get matches for competition %v due to: %w", competition.Name, err)
		}

		// Check if all matches are finished
		finishedCount := 0
		for _, match := range matches {
			if match.Status == repository.UnitStatusCompleted {
				finishedCount += 1
			}
		}
		matchAmount := len(matches)
		if finishedCount == matchAmount {
			// End the competition if there is only one match left
			if matchAmount == 1 {
				_, err := repo.SetCompetitionStatus(ctx, repository.SetCompetitionStatusParams{
					ID:     competition.ID,
					Status: repository.UnitStatusCompleted,
				})
				if err != nil {
					return fmt.Errorf("Failed to finish competition %v due to: %w", competition.Name, err)
				}
			} else {
				// Otherwise generate a new match set
				_, err := GenerateMatchesFromFinishedOnes(ctx, repo, matches)
				if err != nil {
					return fmt.Errorf("Failed to generate a new match set from finished ones for %s due to: %w", competition.ID, err)
				}

			}
		}
	}

	// Check for currently running matches

	// Check if all rounds in them are finished

	// Check if min_rounds is satisfied or if a tie break is required

	// Spawn a new round if needed

	// Or... Set the winner and finish the match

	return nil
}

func IsAnAccpetablePowerOfTwo(n int) bool {
	if n <= 1 { // We only want numbers bigger than or equal to 2
		return false
	}
	// Returns true if only one bit is set (power of two), false otherwise
	return (n & (n - 1)) == 0
}

func GenerateInitialMatches(ctx context.Context, repo *repository.Queries, competition uuid.UUID) ([]repository.Match, error) {
	matches := []repository.Match{}

	inscriptions, err := repo.GetActiveCompetitionInscriptions(ctx, repository.GetActiveCompetitionInscriptionsParams{
		Competition: competition,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to get paricipants: %w", err)
	}

	lenInscriptions := len(inscriptions)
	log.Printf("len(inscriptions) = %v", lenInscriptions)
	if !IsAnAccpetablePowerOfTwo(lenInscriptions) {
		return nil, fmt.Errorf("The number of active participants (%v) is not a power of 2", lenInscriptions)
	}

	step := 2
	for i := 0; i < lenInscriptions; i += step {
		insc1 := inscriptions[i]
		insc2 := inscriptions[i+1]

		match, err := repo.InsertMatch(ctx, repository.InsertMatchParams{
			Competition: competition,
			User1:       insc1.UserID,
			User2:       &insc2.UserID,
			Next:        nil, // TODO: fix me
		})
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}

	return matches, nil
}

func GenerateMatchesFromFinishedOnes(ctx context.Context, repo *repository.Queries, matches []repository.Match) ([]repository.Match, error) {
	newMatches := []repository.Match{}

	lenMatches := len(matches)
	log.Printf("len(matches) = %v", lenMatches)

	step := 2
	for i := 0; i < lenMatches; i += step {
		match1 := matches[i]
		match2 := matches[i+2]

		match, err := repo.InsertMatch(ctx, repository.InsertMatchParams{
			Competition: match1.Competition,
			User1:       *match1.Winner,
			User2:       match2.Winner,
			Next:        nil,
		})
		if err != nil {
			return nil, err
		}

		// Update the next field of the previous matches
		newMatches = append(newMatches, match)
		for _, toUpdate := range []repository.Match{
			match1, match2,
		} {

			_, err := repo.SetNextForMatch(ctx, repository.SetNextForMatchParams{
				ID:   toUpdate.ID,
				Next: &match.ID,
			})
			if err != nil {
				return nil, fmt.Errorf("Failed to update the next field of match %v due to: %w", toUpdate.ID, err)
			}
		}
	}

	return newMatches, nil
}
