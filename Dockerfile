FROM golang:1.13.1 AS builder

WORKDIR /app
COPY . .

RUN go build -o goddns main.go

FROM alpine

ENV CF_API_KEY=xxx \
  CF_API_EMAIL=xxx \
  GODDNS_DOMAINS=xxx \
  DRYRUN=true

WORKDIR /app
COPY --from=builder /app/goddns .
USER golang

CMD ["/app/goddns"]