package main

import (
	"context"
	"fmt"
	"log"
	_http "net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
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
	godotenv.Load()
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	authMid := middleware.NewAuthMiddleware(os.Getenv("AUTH_HOST"), _http.DefaultClient)

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

	nctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	fmt.Println(os.Getenv("PARTY_QUERY_HOST"))

	grcConn, err := grpc.DialContext(nctx, os.Getenv("PARTY_QUERY_HOST"), grpc.WithInsecure())

	if err != nil {
		log.Panicln(err)
	}

	repo := repository.New(r, nc, party.NewPartyQueryServiceClient(grcConn))

	usc := usecase.New(repo)

	log.Fatalln(http.NewHTTP(usc, authMid))
}
