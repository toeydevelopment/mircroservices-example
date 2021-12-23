FROM rust:1.56 as build

WORKDIR /app

# copy over your manifests
COPY . .

RUN cargo build

CMD [ "cargo","run" ]
