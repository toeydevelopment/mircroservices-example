package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/dto"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/pb/party"
)

type IRepository interface {
	CreateParty(ctx context.Context, data dto.CreatePartyDTO) (string, error)
	UpdateParty(ctx context.Context, id string, data dto.UpdatePartyDTO) error
	DeleteParty(ctx context.Context, id string) error
	JoinParty(ctx context.Context, id string, userEmail string) error
	UnJoinParty(ctx context.Context, id string, userEmail string) error
	FindPartyByPartyID(ctx context.Context, id string) (*dto.PartyDTO, error)
}

type Repository struct {
	r          *redis.Client
	nc         *nats.Conn
	partyQuery party.PartyQueryServiceClient
}

func New(
	r *redis.Client,
	nc *nats.Conn,
	partyQuery party.PartyQueryServiceClient,
) IRepository {
	return Repository{r, nc, partyQuery}
}

func (r Repository) CreateParty(ctx context.Context, data dto.CreatePartyDTO) (string, error) {

	// replyID := strconv.Itoa(int(time.Now().Unix()))

	type CreatePartyRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		SeatLimit   int64  `json:"seat_limit"`
		UserEmail   string `json:"user_email"`
		CreatedAt   int64  `json:"created_at"`
	}

	req := CreatePartyRequest{
		Name:        data.Name,
		Description: data.Description,
		SeatLimit:   data.SeatLimit,
		UserEmail:   data.UserEmail,
		CreatedAt:   time.Now().UnixNano() / 1000000,
	}

	b, err := json.Marshal(req)

	if err != nil {
		return "", err
	}

	msg, err := r.nc.RequestMsgWithContext(ctx, &nats.Msg{
		Subject: "party.create",
		Data:    b,
	})

	if err != nil {
		return "", err
	}

	// party id
	return string(msg.Data), nil
}

func (r Repository) UpdateParty(ctx context.Context, id string, data dto.UpdatePartyDTO) error {
	type UpdatePartyRequest struct {
		ID          string  `json:"id"`
		Name        *string `json:"name"`
		Description *string `json:"description"`
		SeatLimit   *int64  `json:"seat_limit"`
		UpdatedAt   int64   `json:"updated_at"`
	}

	req := UpdatePartyRequest{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
		SeatLimit:   data.SeatLimit,
		UpdatedAt:   time.Now().UnixNano() / 1000000,
	}

	b, err := json.Marshal(req)

	if err != nil {
		return err
	}

	_, err = r.nc.RequestMsgWithContext(ctx, &nats.Msg{
		Subject: "party.update",
		Data:    b,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteParty(ctx context.Context, id string) error {
	_, err := r.nc.RequestMsgWithContext(ctx, &nats.Msg{
		Subject: "party.delete",
		Data:    []byte(id),
	})

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) JoinParty(ctx context.Context, id string, userEmail string) error {

	type JoinRequest struct {
		ID        string `json:"id"`
		UserEmail string `json:"user_email"`
	}

	req := JoinRequest{
		ID:        id,
		UserEmail: userEmail,
	}

	b, err := json.Marshal(req)

	if err != nil {
		return err
	}
	_, err = r.nc.RequestMsgWithContext(ctx, &nats.Msg{
		Subject: "party.join",
		Data:    b,
	})

	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UnJoinParty(ctx context.Context, id string, userEmail string) error {

	type UnJoinRequest struct {
		ID        string `json:"id"`
		UserEmail string `json:"user_email"`
	}

	req := UnJoinRequest{
		ID:        id,
		UserEmail: userEmail,
	}

	b, err := json.Marshal(req)

	if err != nil {
		return err
	}
	_, err = r.nc.RequestMsgWithContext(ctx, &nats.Msg{
		Subject: "party.unjoin",
		Data:    b,
	})

	if err != nil {
		return err
	}
	return nil
}

func (r Repository) FindPartyByPartyID(ctx context.Context, id string) (*dto.PartyDTO, error) {

	result, err := r.partyQuery.PartyByID(ctx, &party.PartyByIDRequest{Id: id})

	if err != nil {
		return nil, err
	}

	data := result.GetParty()

	var updatedAt time.Time

	if data.UpdatedAt != nil {
		updatedAt = data.UpdatedAt.AsTime()
	}

	return &dto.PartyDTO{
		ID:          data.GetId(),
		Name:        data.GetName(),
		Description: data.Description,
		SeatLimit:   data.SeatLimit,
		Seat:        data.Seat,
		ImagePath:   data.ImagePath,
		Joined:      data.GetJoined(),
		Owner:       data.GetOwner(),
		CreatedAt:   data.CreatedAt.AsTime(),
		UpdatedAt:   &updatedAt,
	}, nil
}
