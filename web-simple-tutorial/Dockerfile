# Compile stage
FROM registry.cn-hangzhou.aliyuncs.com/shalousun/alpine-goland:1.15.5 AS build-env

# Set proxy
ENV GOPROXY https://goproxy.cn

ADD . /web-simple-tutorial
WORKDIR /web-simple-tutorial

## Add this go mod download command to pull in any dependencies
RUN go mod download

RUN go build -o /app/web-app

# Final stage
FROM registry.cn-hangzhou.aliyuncs.com/shalousun/alpine-zh:3.11

EXPOSE 8080

WORKDIR /
COPY --from=build-env /app/web-app /app/web-app
# copy config files
COPY ./conf/conf.yaml /conf/conf.yaml
# Command to run
CMD ["/app/web-app"]