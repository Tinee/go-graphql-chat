package inmemory

import (
	"sync"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
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
