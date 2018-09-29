package domain

type Storage interface {
	UserRepository() UserRepository
	MessageRepository() MessageRepository
	ProfileRepository() ProfileRepository
}

type UserRepository interface {
	Create(User) (User, error)
	Find(id string) (*User, error)
	Authenticate(username, password string) (*User, error)
}

type MessageRepository interface {
	Create(Message) (Message, error)
}

type ProfileRepository interface {
	Create(Profile) (Profile, error)
}
