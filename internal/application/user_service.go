package application

import "github.com/fabiomattes2016/go-crud/internal/domain"

type UserService struct {
	UserRepository domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.UserRepository.Create(user)
}

func (s *UserService) GetAllUsers() ([]domain.UserResponse, error) {
	users, err := s.UserRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var userResponses []domain.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, domain.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return userResponses, nil
}

func (s *UserService) GetUserByID(id uint) (*domain.UserResponse, error) {
	user, err := s.UserRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &domain.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.UserRepository.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepository.Delete(id)
}
