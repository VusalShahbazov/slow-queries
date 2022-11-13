FROM golang:1.16-alpine as builder-stage

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/main /app/cmd/slow_queries/main.go

FROM alpine as runtime

COPY --from=builder-stage /bin/main .

ENTRYPOINT ["./main"]