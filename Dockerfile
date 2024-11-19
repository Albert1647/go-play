FROM golang:1.23-alpine AS cache

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download


FROM golang:1.23-alpine AS build

WORKDIR /go/src/app

COPY --from=cache /go/pkg /go/pkg
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=/go/bin/app ./


FROM gcr.io/distroless/static-debian12:latest-amd64

COPY --from=build /go/src/app/configs /configs
COPY --from=build /go/bin/app /

CMD [ "/app" ]
