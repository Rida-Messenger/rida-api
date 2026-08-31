ARG GO_VERSION=1.26.7
ARG GO_OS=linux

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN BUILD_FLAGS=""; \
    if ["${APP_ENV}" = "prod"]; then \
        BUILD_FLAGS="-ldflags=-s -w"; \
    fi; \
    CGO_ENABLED=0 GOOS="$GO_OS" \
        go build \
        -trimpath \
        $BUILD_FLAGS \
        -o /out/api \
        ./cmd/api
    
FROM alpine:3.24 as runtime

RUN apk add --no-cache ca-certificates \
    && addgroup -S runner \
    && adduser -S -G runner runner

WORKDIR /app

COPY --from=builder /out/api ./api

USER runner

EXPOSE 3000

CMD [ "./api" ]