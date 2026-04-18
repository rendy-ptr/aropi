FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY apps/backend/go.mod apps/backend/go.sum ./
RUN go mod download

COPY apps/backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o aropi-api ./cmd/main.go

FROM alpine:latest AS runner
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/aropi-api .

EXPOSE 3000
CMD ["./aropi-api"]
