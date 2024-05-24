package service

import (
	"github.com/dMuto/umhelp-servidor-distopico/config"
	"github.com/dMuto/umhelp-servidor-distopico/repo"
	"github.com/rs/zerolog"
)

type Service struct {
	User   *UserService
	Wallet *WalletService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	walletService := newWalletService(cfg, logger, repo)
	userService := NewUserService(cfg, logger, repo, walletService)

	return &Service{
		User:   userService,
		Wallet: walletService,
	}
}
