FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /learn-control ./cmd/learn_control

FROM alpine:latest

COPY --from=builder /learn-control /learn-control

EXPOSE 8000 50051

ENTRYPOINT ["/learn-control"]
