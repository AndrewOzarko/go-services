FROM golang:1.16-alpine as builder

WORKDIR /todo-list

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd/ cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -installsuffix cgo -o /todo-list ./cmd/todo-list

FROM alpine

COPY --from=builder /todo-list/todo-list /todo-list

ENTRYPOINT ["/todo-list"]