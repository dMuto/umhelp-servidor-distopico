package service

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/config"
	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/dMuto/umhelp-servidor-distopico/presenter/req"
	"github.com/dMuto/umhelp-servidor-distopico/repo"
	"github.com/rs/zerolog"
)

type UserService struct {
	Config *config.Config
	Logger *zerolog.Logger
	Repo   *repo.RepoManager
}

func NewUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserService {
	return &UserService{
		Config: cfg,
		Logger: logger,
		Repo:   repo,
	}
}

func (s *UserService) NewUser(ctx context.Context, r *req.User) error {
	user := &model.User{
		Name:     r.Name,
		LastName: r.LastName,
		Document: r.Document,
		Email:    r.Email,
		Password: r.Password,
	}
	s.Logger.Info().Msgf("creating user %s", user.Document)

	return s.Repo.MySQL.User.InsertUser(ctx, user)
}
