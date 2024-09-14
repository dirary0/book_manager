package service

import (
	"book_manager/internal/dto"
	"book_manager/internal/models"
	"book_manager/internal/repository"
	"context"
	"errors"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.UserDTO) error
	UpdateUserByID(ctx context.Context, userID uint, req *dto.UserDTO) error
	DeleteUserByID(ctx context.Context, userID uint) error
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
	ListAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUserByUsername(ctx context.Context, username string, req *dto.UserDTO) error
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) CreateUser(ctx context.Context, req *dto.UserDTO) error {
	user := &models.User{
		Username: req.Username,
	}
	if err := s.userRepository.Create(ctx, user); err != nil {
		return errors.New("创建用户失败: " + err.Error())
	}
	return nil
}

func (s *userService) UpdateUserByID(ctx context.Context, userID uint, req *dto.UserDTO) error {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return errors.New("未找到用户: " + err.Error())
	}
	user.Username = req.Username
	if err := s.userRepository.Update(ctx, user); err != nil {
		return errors.New("更新用户失败: " + err.Error())
	}
	return nil
}

func (s *userService) DeleteUserByID(ctx context.Context, userID uint) error {
	if err := s.userRepository.Delete(ctx, userID); err != nil {
		return errors.New("删除用户失败: " + err.Error())
	}
	return nil
}

func (s *userService) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("未找到用户: " + err.Error())
	}
	return user, nil
}

func (s *userService) ListAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := s.userRepository.AllUsers(ctx)
	if err != nil {
		return nil, errors.New("获取用户列表失败: " + err.Error())
	}
	return users, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("未找到用户: " + err.Error())
	}
	return user, nil
}

func (s *userService) UpdateUserByUsername(ctx context.Context, username string, req *dto.UserDTO) error {
	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return errors.New("未找到用户: " + err.Error())
	}

	user.Username = req.Username

	if err := s.userRepository.UpdateByUsername(ctx, username, user); err != nil {
		return errors.New("更新用户失败: " + err.Error())
	}
	return nil
}
