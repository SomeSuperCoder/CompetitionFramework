package matchmaking

import (
	"context"
	"fmt"
	"log"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
	"github.com/google/uuid"
)

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
			Next:        nil,
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
		match2 := matches[i+1]

		match, err := repo.InsertMatch(ctx, repository.InsertMatchParams{
			Competition: match1.Competition,
			User1:       *match1.Winner,
			User2:       match2.Winner,
			Next:        nil,
		})
		if err != nil {
			return nil, fmt.Errorf("Failed to insert match due to: %w", err)
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
