### build stage
FROM golang:1.18.3-alpine AS build-env
RUN apk --no-cache add build-base git curl
ADD . /build
WORKDIR /build/app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o balance

# delete main.go
RUN rm *.go

### final stage
# why not scratch? We need bash to view logs
FROM alpine:latest

COPY --from=build-env /build/app/cmd/api/ /app/

WORKDIR /app

ENTRYPOINT ["./balance"]