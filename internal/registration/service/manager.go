package service

import (
	"xamss.onelab.final/internal/config"
	"xamss.onelab.final/internal/registration/repository"
	"xamss.onelab.final/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.JWTToken
	Config     *config.Config
}

func New(repository repository.Repository, token *jwttoken.JWTToken, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     config,
	}
}
