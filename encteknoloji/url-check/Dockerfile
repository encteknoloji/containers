FROM golang:1.19-alpine as builder

WORKDIR /app

COPY check_url.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" check_url.go \
  && mkdir -p /app/etc/ssl/certs/ \
  && cp /etc/ssl/certs/ca-certificates.crt /app/etc/ssl/certs/
COPY passwd /app/etc/passwd

FROM scratch

COPY --from=builder /app/ /

USER nobody

CMD ["/check_url"]