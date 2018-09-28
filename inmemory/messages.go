package inmemory

import (
	"sync"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
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
