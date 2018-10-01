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

func (m *userInMemory) Create(u domain.User) (domain.User, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, ErrCreatingUser
	}
	// Adds data to the input user.
	u.Password = string(bs)
	u.ID = generateID()
	u.CreatedAt = time.Now()

	m.mtx.Lock()
	defer m.mtx.Unlock()
	// if user exists bail.
	for _, v := range m.users {
		if v.Username == u.Username {
			return domain.User{}, ErrUserExists
		}
	}
	m.users = append(m.users, u)

	return u, nil
}

// Find tries to find the user by ID.
func (m *userInMemory) Find(id string) (*domain.User, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, v := range m.users {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, ErrUserNotFound
}

func (m *userInMemory) Authenticate(username, password string) (*domain.User, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, v := range m.users {
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
