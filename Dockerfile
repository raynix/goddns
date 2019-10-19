FROM golang:1.13.1 AS builder

ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app
COPY . .
RUN go build -a -installsuffix cgo -o goddns main.go && \
  addgroup --system --gid 2000 golang && \
  adduser --system --gid 2000 --uid 2000 golang

FROM scratch

ENV CF_API_KEY=xxx \
  CF_API_EMAIL=xxx \
  GODDNS_DOMAINS=xxx \
  DRYRUN=true

WORKDIR /app
COPY --from=builder /app/goddns .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER golang

ENTRYPOINT ["/app/goddns"]
