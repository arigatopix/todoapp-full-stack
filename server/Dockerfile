FROM golang:1.18.3-buster AS build

WORKDIR /app

COPY ./server/go.mod ./
COPY ./server/go.sum ./

RUN go mod download

COPY ./server/ ./

RUN go build -o server

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/server ./

CMD ["/app/server"]