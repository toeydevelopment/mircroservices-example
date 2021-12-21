package main

import (
	"context"
	"log"
	_http "net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/delivery/http"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/middleware"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/pb/party"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/repository"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/usecase"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	authMid := middleware.NewAuthMiddleware("", _http.DefaultClient)

	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Panicln(err)
	}

	defer nc.Close()

	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		log.Panicln(err)
	}

	defer r.Close()

	nctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	// grpc.WithBlock()
	grcConn, err := grpc.DialContext(nctx, "", grpc.WithInsecure())

	if err != nil {
		log.Panicln(err)
	}

	repo := repository.New(r, nc, party.NewPartyQueryServiceClient(grcConn))

	usc := usecase.New(repo)

	log.Fatalln(http.NewHTTP(usc, authMid))
}
