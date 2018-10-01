package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
)

var (
	ErrMessageNotFound = errors.New("message not found in memory")
)

type messagesInMemory struct {
	mtx      *sync.Mutex
	messages []domain.Message
}

func (m *messagesInMemory) Create(ms domain.Message) (domain.Message, error) {
	ms.ID = generateID()
	ms.CreatedAt = time.Now()

	m.mtx.Lock()
	m.messages = append(m.messages, ms)
	m.mtx.Unlock()

	return ms, nil
}

func (m *messagesInMemory) Find(id string) (*domain.Message, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, m := range m.messages {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, ErrMessageNotFound
}
