package domain

type Storage interface {
	UserRepository() UserRepository
}

type UserRepository interface {
	Create(User) (User, error)
	Find(id string) (*User, error)
	Authenticate(username, password string) (*User, error)
}
