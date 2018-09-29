package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
)

var (
	ErrProfileNotFound = errors.New("profile not found in memory")
)

type profileInMemory struct {
	mtx      *sync.Mutex
	profiles []domain.Profile
}

func (m *profileInMemory) Create(p domain.Profile) (domain.Profile, error) {
	p.ID = generateID()
	p.CreatedAt = time.Now()

	m.mtx.Lock()
	m.profiles = append(m.profiles, p)
	m.mtx.Unlock()

	return p, nil
}

func (m *profileInMemory) Find(id string) (*domain.Profile, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, v := range m.profiles {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, ErrUserNotFound
}
