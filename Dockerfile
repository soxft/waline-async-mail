FROM golang:1.20-alpine as builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN set -evx -o pipefail        \
    && apk update               \
    && apk add --no-cache git   \
    && rm -rf /var/cache/apk/*  \
    && go build -ldflags="-s -w" -o waline_mailer main.go

FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/waline_mailer /app/waline_mailer
COPY --from=builder /app/config.yaml /app/config.yaml

CMD ["./waline_mailer", "-c", "config.yaml"]