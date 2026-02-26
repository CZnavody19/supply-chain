FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . ./
RUN CGO_ENABLED=0 go build -o main .

FROM scratch
WORKDIR /app
COPY configuration/ ./configuration/
COPY migrations/ ./migrations/
COPY --from=builder /app/main .
CMD ["./main"]
