package usecase

import (
	"errors"

	"github.com/ryatsuga/codepix-go/domain/model"
)

type PixUseCase struct {
	PixRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	p.PixRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("Unable to register key")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
