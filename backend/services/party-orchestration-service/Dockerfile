FROM golang:1.17-alpine3.14 as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go


FROM alpine:3.14

WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /build/main .

EXPOSE 3000

EXPOSE 50051

CMD ["/app/main"]
