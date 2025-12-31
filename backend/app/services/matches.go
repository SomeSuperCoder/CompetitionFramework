package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SomeSuperCoder/CompetitionFramework/backend/repository"
)

type MatchService struct {
	Repo *repository.Queries
}

func IsAnAccpetablePowerOfTwo(n int) bool {
	if n <= 1 { // We only want numbers bigger than or equal to 2
		return false
	}
	// Returns true if only one bit is set (power of two), false otherwise
	return (n & (n - 1)) == 0
}

func (s *MatchService) Generate(r *http.Request, args *repository.GetActiveCompetitionInscriptionsParams, reply *[]repository.Match) error {
	matches := []repository.Match{}

	inscriptions, err := s.Repo.GetActiveCompetitionInscriptions(r.Context(), *args)
	if err != nil {
		return fmt.Errorf("Failed to get paricipants: %w", err)
	}

	lenInscriptions := len(inscriptions)
	fmt.Printf("len(inscriptions) = %v\n", lenInscriptions)
	if !IsAnAccpetablePowerOfTwo(lenInscriptions) {
		return fmt.Errorf("The number of active participants (%v) is not a power of 2", lenInscriptions)
	}

	step := 2
	for i := 0; i < lenInscriptions; i += step {
		fmt.Println("we are here")
		insc1 := inscriptions[i]
		insc2 := inscriptions[i+1]

		fmt.Println(insc1.ID)
		fmt.Println(insc2.ID)
		match, err := s.Repo.InsertMatch(r.Context(), repository.InsertMatchParams{
			Competition: args.Competition,
			StartTime:   time.Now(),                    // TODO: move to rounds
			EndTime:     time.Now().Add(time.Hour * 2), // TODO: fix me
			User1:       insc1.UserID,
			User2:       &insc2.UserID,
			Prev:        nil, // TODO: fix me
		})
		if err != nil {
			return err
		}
		matches = append(matches, match)
	}

	*reply = matches
	return nil
}
