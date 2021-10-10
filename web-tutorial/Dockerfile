# Compile stage
FROM golang:1.14.6 AS build-env

ADD . /web-tutorial
WORKDIR /web-tutorial

## Add this go mod download command to pull in any dependencies
RUN go mod download

RUN go build -o /server

# Final stage
FROM alpine:3.9

EXPOSE 8080

WORKDIR /
COPY --from=build-env /server /app

CMD ["/app/server","-config=./test.toml"]