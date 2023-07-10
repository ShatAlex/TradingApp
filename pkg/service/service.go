package service

import "github.com/ShatAlex/trading-app/pkg/repository"

type Authorization interface {
}

type Trade interface {
}

type Type interface {
}

type Service struct {
	Authorization
	Trade
	Type
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
