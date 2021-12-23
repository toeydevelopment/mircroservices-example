package grpc

import (
	"context"

	"github.com/toeydevelopment/microservices-example/party-query-service/pb/party"
	"github.com/toeydevelopment/microservices-example/party-query-service/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	usc usecase.IUsecase
	party.UnimplementedPartyQueryServiceServer
}

func newHandler(usc usecase.IUsecase) handler {
	return handler{usc: usc}
}

func (h handler) PartyByID(ctx context.Context, in *party.PartyByIDRequest) (*party.PartyByIDResponse, error) {

	p, err := h.usc.FindPartiyByID(ctx, in.GetId())

	if err != nil {
		return nil, err
	}

	return &party.PartyByIDResponse{
		Party: &party.Party{
			Id:          p.ID.Hex(),
			Name:        p.Name,
			Description: &p.Description,
			SeatLimit:   &p.SeatLimit,
			Seat:        &p.Seat,
			ImagePath:   &p.ImagePath,
			Joined:      p.Joined,
			Owner:       p.Owner,
			CreatedAt:   timestamppb.New(p.CreatedAt),
			UpdatedAt:   timestamppb.New(p.UpdatedAt),
			DeletedAt:   timestamppb.New(p.DeletedAt),
		},
	}, nil
}
