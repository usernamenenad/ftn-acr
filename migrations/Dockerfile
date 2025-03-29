FROM alpine:latest AS builder


RUN apk update && apk add --no-cache \
go \
make \
gcc \
musl-dev \
libc-dev

RUN go install -v -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:latest

COPY --from=builder /root/go/bin/migrate /usr/local/bin/migrate

COPY . . 

RUN apk update && apk add --no-cache ca-certificates

CMD ["sh", "migrate.sh"]
