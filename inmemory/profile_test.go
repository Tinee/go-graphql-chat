package inmemory_test

import (
	"testing"

	"github.com/Tinee/go-graphql-chat/domain"
)

func Test_profileInMemory_Create(t *testing.T) {
	c := NewClient()
	repo := c.ProfileRepository()

	p, err := repo.Create(domain.Profile{
		Age:       25,
		FirstName: "Foo",
		LastName:  "Bar",
		UserID:    "Test",
	})

	if err != nil {
		t.Errorf("Expected not an error but got: %v", err)
	}
	// hiding the error, because I'm not testing that method.
	pp, _ := repo.Find(p.ID)

	if p.FirstName != pp.FirstName {
		t.Errorf("Expected (%v) to (%v) the same, but they're not", p.FirstName, pp.FirstName)
	}
}

func Test_profileInMemory_FindMany(t *testing.T) {
	c := NewClient()
	c.FillWithMockData()
	repo := c.ProfileRepository()
	t.Parallel()
	type args struct {
		take   int
		offset int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "Should find two items",
			args: args{
				take:   2,
				offset: 2,
			},
			wantCount: 2,
		},
		{
			name: "Should find zero items, because we don't have that many profiles in memory.",
			args: args{
				take:   5,
				offset: 5,
			},
			wantCount: 0,
		},
		{
			name: "Should take the full list if we don't have an offset.",
			args: args{
				take:   10,
				offset: 0,
			},
			wantCount: 4,
		},
		{
			name: "Should only take two items, when we have a small offset.",
			args: args{
				take:   10,
				offset: 2,
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repo.FindMany(tt.args.take, tt.args.offset)

			if len(got) != tt.wantCount {
				t.Errorf("Expected a list count of (%v) but got (%v)", tt.wantCount, len(got))
			}
		})
	}
}
