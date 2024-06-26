package service

import (
	"github.com/dMuto/umhelp-servidor-distopico/config"
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

/*func (s *WalletService) New(tx *sqlx.Tx, ctx context.Context, ) error {
	wallet := &model.Wallet{
		IDOwner:    userId,
		IDCurrency: currencyId,
		Alias:      alias,
	}


	s.Logger.Info().Msgf("creating wallet for user %d", wallet.IDOwner)

	return s.Respo.MySQL.Wallet.InsertWallet(tx ,ctx, wallet)
}*/
