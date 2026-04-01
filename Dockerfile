# ========================
# 1. BUILD STAGE
# ========================
FROM golang:1.26-alpine AS builder

WORKDIR /app

# copy go mod
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build binary (arah ke cmd/main.go)
RUN go build -o main ./cmd/main.go


# ========================
# 2. RUN STAGE (lebih ringan)
# ========================
FROM alpine:latest

WORKDIR /app

# copy binary dari builder
COPY --from=builder /app/main .

# copy certs (WAJIB untuk TLS Aiven)
COPY certs ./certs

# expose port
EXPOSE 8080

# run app
CMD ["./main"]
