package service

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/config"
	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/dMuto/umhelp-servidor-distopico/repo"
	"github.com/rs/zerolog"
)

type WalletService struct {
	Config *config.Config
	Logger *zerolog.Logger
	Respo  *repo.RepoManager
}

func newWalletService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *WalletService {
	return &WalletService{
		Config: cfg,
		Logger: logger,
		Respo:  repo,
	}
}

func (s *WalletService) New(ctx context.Context, userId, currencyId int64, alias string) error {
	wallet := &model.Wallet{
		IDOwner:    userId,
		IDCurrency: currencyId,
		Alias:      alias,
	}

	s.Logger.Info().Msgf("creating wallet for user %d", wallet.IDOwner)

	return s.Respo.MySQL.Wallet.InsertWallet(ctx, wallet)
}
