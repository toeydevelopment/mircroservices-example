package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/toeydevelopment/microservices-example/party-query-service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	FindParties(ctx context.Context, afterID string, limit int64) ([]entity.Party, error)
	FindPartiyByID(ctx context.Context, id string) (*entity.Party, error)
}

type Repository struct {
	nc    *nats.Conn
	redis *redis.Client
	m     *mongo.Client
}

func New(
	nc *nats.Conn,
	redis *redis.Client,
	m *mongo.Client,
) IRepository {
	return Repository{
		nc:    nc,
		redis: redis,
		m:     m,
	}
}

func (r Repository) FindParties(ctx context.Context, afterID string, limit int64) ([]entity.Party, error) {

	opt := options.Find()

	opt.SetLimit(limit)

	coll := r.m.Database("party").Collection("party")

	condition := bson.M{}

	if afterID != "" {

		id, err := primitive.ObjectIDFromHex(afterID)

		if err != nil {
			return nil, err
		}

		condition = bson.M{
			"_id": bson.M{
				"$gt": id,
			},
		}
	}

	cursor, err := coll.Find(ctx, condition, opt)

	if err != nil {
		return nil, err
	}

	var data []entity.Party

	for cursor.TryNext(ctx) {
		var d entity.Party

		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}

func (r Repository) FindPartiyByID(ctx context.Context, id string) (*entity.Party, error) {

	coll := r.m.Database("party").Collection("party")

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var data entity.Party

	if err := coll.FindOne(ctx, bson.M{
		"_id": oid,
	}).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
