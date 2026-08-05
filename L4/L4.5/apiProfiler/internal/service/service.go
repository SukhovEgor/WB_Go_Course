package service

const (
	defaultName  = "User"
	defaultEmail = "user@example.com"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Sum(a, b int) int {
	return a + b
}

func (s *Service) GetUsers() []User {

	users := make([]User, 0, 10000)

	for i := 1; i <= 10000; i++ {

		users = append(users, User{
			ID:    i,
			Name:  defaultName,
			Email: defaultEmail,
		})
	}

	return users
}