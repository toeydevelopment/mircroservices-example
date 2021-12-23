package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/toeydevelopment/microservices-example/party-query-service/delivery/grpc"
	"github.com/toeydevelopment/microservices-example/party-query-service/delivery/http"
	"github.com/toeydevelopment/microservices-example/party-query-service/helper"
	"github.com/toeydevelopment/microservices-example/party-query-service/repository"
	"github.com/toeydevelopment/microservices-example/party-query-service/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	godotenv.Load(".env")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nc, err := nats.Connect(os.Getenv("NATS_HOST"))

	if err != nil {
		log.Panicln(err)
	}

	defer nc.Close()

	r := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		log.Panicln(err)
	}

	defer r.Close()

	c := options.Client().ApplyURI(
		fmt.Sprintf(
			"mongodb://%s:%s@%s",
			os.Getenv("MONGO_USERNAME"),
			os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_HOST"),
		),
	)

	m, err := mongo.Connect(ctx, c)

	if err != nil {
		log.Panicln(err)
	}

	defer m.Disconnect(ctx)

	repo := repository.New(nc, r, m)

	usc := usecase.New(repo, helper.New())

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		log.Fatalln(grpc.NewGRPCServer("0.0.0.0:50051", usc))
	}()

	wg.Add(1)
	go func() {
		log.Fatalln(http.NewHTTP(usc))
	}()

	wg.Wait()

	log.Println("service exited successfully")
}
