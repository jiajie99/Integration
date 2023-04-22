package go_snaps

import (
	"testing"
	"time"

	"github.com/gkampitakis/go-snaps/match"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/gofrs/uuid"
)

type data struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func generateUUIDV4() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}

func TestGoSnaps(t *testing.T) {
	tests := []struct {
		name string
		data *data
	}{
		{
			name: "case 1",
			data: &data{
				ID:        generateUUIDV4(),
				Name:      "Jacky",
				CreatedAt: time.Now(),
			},
		},
		{
			name: "case 2",
			data: &data{
				ID:        generateUUIDV4(),
				Name:      "Tom",
				CreatedAt: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Ignore dynamic data.
			snaps.MatchJSON(t, tt.data, match.Any("id", "created_at"))
		})
	}
}
