FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.5.3/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migrations

COPY ./migrations/*.sql ./
COPY ./migrations/migration.sh ./

RUN chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]