FROM golang:1.20 AS builder

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /pfsense-http-wol

FROM alpine:3.17
COPY --from=builder /pfsense-http-wol /pfsense-http-wol
EXPOSE 8080

CMD ["/pfsense-http-wol"]
