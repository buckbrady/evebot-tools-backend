FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o zkillsync ./cmd/zkillsync

FROM gcr.io/distroless/static:nonroot as app

LABEL org.opencontainers.image.source="https://github.com/buckbrady/evebot-tools-backend"

USER nonroot:nonroot

COPY --from=builder --chown=nonroot:nonroot /app/zkillsync /zkillsync

ENTRYPOINT ["/zkillsync"]