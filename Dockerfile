FROM golang:1.17.1-alpine3.14 AS builder
WORKDIR /app
COPY ./src ./src
COPY main.go ./
COPY go.mod ./
COPY go.sum ./
RUN apk add --update gcc musl-dev
RUN go mod download
RUN go build -o /app.out

FROM golang:1.17.1-alpine3.14 AS production
WORKDIR /
COPY --from=builder /app.out ./
EXPOSE 8000
CMD [ "/app.out" ]

FROM golang:1.17.1 AS development
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /app
EXPOSE 8000
CMD [ "air" ]

