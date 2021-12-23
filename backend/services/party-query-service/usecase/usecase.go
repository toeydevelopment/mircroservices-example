package usecase

import (
	"context"

	"github.com/toeydevelopment/microservices-example/party-query-service/entity"
	"github.com/toeydevelopment/microservices-example/party-query-service/helper"
	"github.com/toeydevelopment/microservices-example/party-query-service/repository"
)

type IUsecase interface {
	FindParties(ctx context.Context, afterID string, limit int64) ([]entity.Party, string, error)
	FindPartiyByID(ctx context.Context, id string) (*entity.Party, error)
}

type Usecase struct {
	repo repository.IRepository
	h    helper.IHelper
}

func New(repo repository.IRepository, h helper.IHelper) IUsecase {
	return Usecase{repo, h}
}

func (u Usecase) FindParties(ctx context.Context, cursor string, limit int64) ([]entity.Party, string, error) {

	if limit == 0 {
		limit = 10
	}

	if limit > 20 {
		limit = 20
	}

	afterID, _ := u.h.DecodeCursor(cursor)

	parties, err := u.repo.FindParties(ctx, afterID, limit)

	if err != nil {
		return nil, "", err
	}

	last := parties[len(parties)-1].ID.Hex()

	newCursor := u.h.EncodeCursor(last)

	return parties, newCursor, nil
}

func (u Usecase) FindPartiyByID(ctx context.Context, id string) (*entity.Party, error) {
	party, err := u.repo.FindPartiyByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return party, nil
}
