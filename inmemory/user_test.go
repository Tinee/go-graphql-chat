package inmemory_test

import (
	"testing"

	"github.com/Tinee/go-graphql-chat/domain"
	"github.com/Tinee/go-graphql-chat/inmemory"
)

func Test_userInMemory_Create(t *testing.T) {
	c := NewClient()
	repo := c.UserRepository()
	new := domain.User{
		Username: "Foo",
		Password: "Bar",
	}
	u, err := repo.Create(new)

	if err != nil {
		t.Errorf("Expected not an error but got: %v", err)
	}
	// hiding the error, because I'm not testing that method.
	cu, _ := repo.Find(u.ID)

	if u.Username != cu.Username {
		t.Errorf("Expected (%v) to (%v) to be the same, but they're not", u.Username, cu.Username)
	}

	if u.Password != cu.Password {
		t.Errorf("Expected (%v) to (%v) to be the same, but they're not", u.Password, cu.Password)
	}

	if new.Password == cu.Password {
		t.Errorf("Expected (%v) to (%v) not be the same, because of password hashing but they're are.", u.Password, cu.Password)
	}
	// Should throw an error, because we're trying to insert someone with the same Username.
	_, err = repo.Create(new)
	if err != inmemory.ErrUserExists {
		t.Errorf("Expected err to be (%v) but got (%v)", inmemory.ErrUserExists, err)
	}
}

func Test_userInMemory_Authenticate(t *testing.T) {
	c := NewClient()
	c.FillWithMockData()
	repo := c.UserRepository()

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should find user, and by finding the user it should give back nil as an error.",
			args: args{
				username: "tine",
				password: "test1",
			},
			wantErr: false,
		},
		{
			name: "Should find the user in memory, BUT should fail because it doesn't have the same password as the mock.",
			args: args{
				username: "tine",
				password: "notValid",
			},
			wantErr: true,
		},
		{
			name: "Should not find the user in memory, which should then throw an error",
			args: args{
				username: "fooIsNotInMemory",
				password: "foHasNoPassword",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userInMemory.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
