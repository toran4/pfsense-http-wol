from golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /pfsense-http-wol
EXPOSE 8080

CMD ["/pfsense-http-wol"]