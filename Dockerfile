### build stage
FROM golang:1.18.3-alpine AS build-env
RUN apk --no-cache add build-base git curl
ADD . /build
WORKDIR /build/app/cmd/api

RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' -o=balance

### final stage
# why not scratch? We need bash to connect to the conteiner
FROM alpine:latest

COPY --from=build-env /build/app/cmd/api/balance /app/

WORKDIR /app

ENTRYPOINT ["./balance"]