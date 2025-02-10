ARG GOLANG_VERSION=1.22.4
FROM golang:${GOLANG_VERSION} as builder

WORKDIR /app

ADD ../go.mod go.sum /modules/

RUN --mount=type=cache,target=/tmp/go_cache/ \
    update-ca-certificates && cd /modules && go mod download

COPY .. ./
RUN --mount=type=cache,target=/tmp/go_cache/ \
    CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd

COPY /config/local.json /out/config/local.json

FROM alpine:3.13 AS build-release-stage

ARG EXPOSE_PORT

WORKDIR /opt/microservice
COPY --from=builder /out/ /opt/microservice/

#Goose migrations
ADD https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

COPY migrations/* /opt/microservice/migrations/
COPY ./config/local.json /opt/microservice/config/

ADD docker/migrations.sh /opt/microservice/migrations.sh
RUN chmod +x /opt/microservice/migrations.sh

EXPOSE ${EXPOSE_PORT}

ENTRYPOINT ["/opt/microservice/app"]
