package server

import (
	"fmt"
	"os"
)

type BackendRepository interface {
	GetMessageByName(name string) (string, error)
}

func NewUserRepository() BackendRepository {
	return &userRepository{}
}

type userRepository struct{}

func (u *userRepository) GetMessageByName(name string) (string, error) {
	switch name != "" {
	case true:
		return fmt.Sprintf("Hello, %v, Nice to meet you!, Backend IP: %v ", name, os.Getenv("MY_POD_IP")), nil
	default:
		return "", fmt.Errorf("no name error")
	}
}
