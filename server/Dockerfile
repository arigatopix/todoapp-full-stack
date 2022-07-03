FROM golang:1.18.3-buster AS build

WORKDIR /app

COPY ./server/go.mod ./
COPY ./server/go.sum ./

RUN go mod download

COPY ./server/ ./

RUN go build -o server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/server ./

CMD ["/server"]