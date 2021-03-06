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
	Find(string) (*Message, error)
}

type ProfileRepository interface {
	Create(Profile) (Profile, error)
	Find(id string) (*Profile, error)
	FindMany(take int, offset int) []Profile
}
