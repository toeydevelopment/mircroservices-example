package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/toeydevelopment/microservices-example/party-orchestration-service/dto"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/repository"
)

var (
	ErrOnlyOwner         = errors.New("only party owners are allowed")
	ErrNotFound          = errors.New("not found party")
	ErrAlreadyJoined     = errors.New("already joined party")
	ErrOperationConflict = errors.New("conflict not exists in party")
)

type IUsecase interface {
	CreateParty(ctx context.Context, data dto.CreatePartyDTO) (string, error)
	UpdateParty(ctx context.Context, id string, data dto.UpdatePartyDTO) error
	DeleteParty(ctx context.Context, id, userEmail string) error
	JoinParty(ctx context.Context, id, userEmail string) error
	UnJoinParty(ctx context.Context, id, userEmail string) error
}

type Usecase struct {
	repo repository.IRepository
}

func New(
	repo repository.IRepository,
) IUsecase {
	return Usecase{repo}
}

func (u Usecase) CreateParty(ctx context.Context, data dto.CreatePartyDTO) (string, error) {

	id, err := u.repo.CreateParty(ctx, data)

	if err != nil {
		log.Println("create party failed error: ", err)
		return "", err
	}

	return id, nil
}

func (u Usecase) UpdateParty(ctx context.Context, id string, data dto.UpdatePartyDTO) error {
	party, err := u.repo.FindPartyByPartyID(ctx, id)

	if err != nil {
		log.Println("find party failed error: ", err)
		return err
	}

	if !isOwner(party.Owner, data.UserEmail) {
		return ErrOnlyOwner
	}

	if err := u.repo.UpdateParty(ctx, id, data); err != nil {
		log.Println("update party failed error: ", err)

	}

	return nil
}

func (u Usecase) DeleteParty(ctx context.Context, id, userEmail string) error {
	party, err := u.repo.FindPartyByPartyID(ctx, id)

	if err != nil {
		log.Println("find party failed error: ", err)
		return err
	}

	if !isOwner(party.Owner, userEmail) {
		return ErrOnlyOwner
	}

	if err := u.repo.DeleteParty(ctx, id); err != nil {
		log.Println("delete party failed error: ", err)

	}

	return nil
}

func (u Usecase) JoinParty(ctx context.Context, id, userEmail string) error {
	party, err := u.repo.FindPartyByPartyID(ctx, id)

	if err != nil {
		log.Println("find party failed error: ", err)
		return err
	}

	if existsInParty(party.Joined, userEmail) {
		return ErrAlreadyJoined
	}

	if err := u.repo.JoinParty(ctx, id, userEmail); err != nil {
		log.Println("join party failed error: ", err)
		return err
	}

	return nil
}

func (u Usecase) UnJoinParty(ctx context.Context, id, userEmail string) error {

	party, err := u.repo.FindPartyByPartyID(ctx, id)

	if err != nil {
		log.Println("find party failed error: ", err)
		return err
	}

	if !existsInParty(party.Joined, userEmail) {
		return ErrOperationConflict
	}

	if err := u.repo.UnJoinParty(ctx, id, userEmail); err != nil {
		log.Println("umjoin party failed error: ", err)
		return err
	}

	return nil
}

func isOwner(owner string, userEmail string) bool {
	return owner == userEmail
}

func existsInParty(joiners []string, email string) bool {
	joined := false

	for _, party := range joiners {
		if party == email {
			joined = true
			break
		}
	}

	return joined
}
