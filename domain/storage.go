package domain

type Storage interface {
	UserRepository() UserRepository
	MessageRepository() MessageRepository
}

type UserRepository interface {
	Create(User) (User, error)
	Find(id string) (*User, error)
	Authenticate(username, password string) (*User, error)
}

type MessageRepository interface {
	Create(Message) (Message, error)
}
