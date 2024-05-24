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
	Config        *config.Config
	Logger        *zerolog.Logger
	Repo          *repo.RepoManager
	walletService *WalletService
}

func NewUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager, walletService *WalletService) *UserService {
	return &UserService{
		Config:        cfg,
		Logger:        logger,
		Repo:          repo,
		walletService: walletService,
	}
}

const BRL = 1

func (s *UserService) NewUser(ctx context.Context, r *req.User) (u *model.User, err error) {
	user := &model.User{
		Name:     r.Name,
		LastName: r.LastName,
		Document: r.Document,
		Email:    r.Email,
		Password: r.Password,
	}

	//tx, err := s.Repo.MySQL.Beg

	userId, err := s.Repo.MySQL.User.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	err = s.walletService.New(ctx, userId, BRL, "Default Wallet")
	if err != nil {
		return err
	}

	return nil
}
