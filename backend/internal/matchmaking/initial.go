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
	log.Println("Generating inital matches!")
	matches := []repository.Match{}

	inscriptions, err := repo.GetCompetitionInscriptions(ctx, repository.GetCompetitionInscriptionsParams{
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
