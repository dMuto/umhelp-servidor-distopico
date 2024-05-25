package service

import (
	"context"
	"fmt"
	"strings"

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
}

func NewUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager, walletService *WalletService) *UserService {
	return &UserService{
		Config:        cfg,
		Logger:        logger,
		Repo:          repo,
	}
}

func (s *UserService) NewUser(ctx context.Context, r *req.User) (u *model.User, err error) {
	user := &model.User{
		Name:     r.Name,
		LastName: r.LastName,
		Document: r.Document,
		Email:    r.Email,
		Password: r.Password,
	}

	tx, err := s.Repo.MySQL.BeginReadCommittedTx(ctx)
	if err != nil{
		return nil, err
	}

	defer tx.Rollback()

	userId, err := s.Repo.MySQL.User.InsertUser(tx, ctx, user)
	if err != nil {
		return nil, err
	}

	user.IdUser = userId

	currency, found, err := s.Repo.MySQL.Currency.SelectByCurrencyCode(tx, ctx, model.CurrencyBRL)
	if err != nil{
		return nil, err
	}

	if !found{
		return nil, fmt.Errorf("cannot find `%s` currency in database", model.CurrencyBRL)
	}
	
	walletModel := &model.Wallet{
		IDOwner: user.IdUser,
		IDCurrency: currency.IDCurrency,
		Alias: strings.Join([]string{user.Name + "'s", "Wallet"}, " "),
	}

	if err = s.Repo.MySQL.Wallet.InsertWallet(tx, ctx, walletModel); err != nil {
		return nil, err
	}

	return user, nil
}
