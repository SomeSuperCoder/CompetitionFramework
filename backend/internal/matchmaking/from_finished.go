package matchmaking

import (
	"context"
	"fmt"
	"log"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

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
