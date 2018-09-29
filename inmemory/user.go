package inmemory

import (
	"errors"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Tinee/go-graphql-chat/domain"
)

// Common errors for users
var (
	ErrUserNotFound = errors.New("user not found in memory")
	ErrCreatingUser = errors.New("couldn't create the user")
	ErrUserExists   = errors.New("couldn't create an existing user")
)

type userInMemory struct {
	mtx   *sync.Mutex
	users []domain.User
}

func (uim *userInMemory) Create(u domain.User) (domain.User, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, ErrCreatingUser
	}
	// Adds data to the input user.
	u.Password = string(bs)
	u.ID = generateID()
	u.CreatedAt = time.Now()

	uim.mtx.Lock()
	defer uim.mtx.Unlock()
	// if user exists bail.
	for _, v := range uim.users {
		if v.Username == u.Username {
			return domain.User{}, ErrUserExists
		}
	}
	uim.users = append(uim.users, u)

	return u, nil
}

// Find tries to find the user by ID.
func (uim *userInMemory) Find(id string) (*domain.User, error) {
	uim.mtx.Lock()
	defer uim.mtx.Unlock()
	for _, v := range uim.users {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, ErrUserNotFound
}

func (uim *userInMemory) Authenticate(username, password string) (*domain.User, error) {
	uim.mtx.Lock()
	defer uim.mtx.Unlock()
	for _, v := range uim.users {
		if v.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(password))
			if err != nil {
				return nil, ErrUserNotFound
			}
			return &v, nil
		}
	}
	return nil, ErrUserNotFound
}
