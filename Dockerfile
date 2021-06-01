# Initial stage: download modules
FROM golang:1.16-alpine as builder
RUN mkdir /build/

ADD . /build/
WORKDIR /build/

ENV config=docker
ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"

RUN go mod vendor

RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o rest-api-kata ./cmd/api/main.go

# Intermediate stage: Build the binary
FROM golang:1.16-alpine as runner

RUN apk add --no-cache ca-certificates supervisor bash tzdata
RUN apk --no-cache add redis \
 && sed -i 's/protected-mode yes/protected-mode no/' /etc/redis.conf \
 && sed -i 's/^\(bind .*\)$/# \1/' /etc/redis.conf \
 && sed -i 's/^\(daemonize .*\)$/# \1/' /etc/redis.conf \
 && sed -i 's/^\(dir .*\)$/# \1\ndir \/data/' /etc/redis.conf \
 && sed -i 's/^\(logfile .*\)$/# \1/' /etc/redis.conf

RUN mkdir -p /etc/supervisor/conf.d/
RUN mkdir -p /rest-api-kata
RUN mkdir -p /rest-api-kata/config

RUN cp /usr/share/zoneinfo/Europe/Istanbul /etc/localtime
RUN echo "Europe/Istanbul" > /etc/timezone

RUN apk del tzdata

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

VOLUME /data

ENV config=docker

COPY --from=builder /build/rest-api-kata /rest-api-kata/rest-api-kata
COPY --from=builder /build/config /rest-api-kata/config

WORKDIR /rest-api-kata

EXPOSE 8080
EXPOSE 6379

HEALTHCHECK CMD redis-cli ping

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]

