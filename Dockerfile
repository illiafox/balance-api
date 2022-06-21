### build stage
FROM golang:1.18.3-alpine AS build-env
RUN apk --no-cache add build-base git curl
ADD . /build
WORKDIR /build/app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o service

# delete main.go
RUN rm *.go

### final stage
FROM scratch
COPY --from=build-env /build/app/cmd/api/ /app/

WORKDIR /app

ENTRYPOINT ["./service"]