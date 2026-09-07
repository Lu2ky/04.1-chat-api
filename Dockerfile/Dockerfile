FROM golang:1.26.7-alpine3.24 AS builder

WORKDIR /api
COPY . .
RUN go mod tidy
RUN go test -v -cover
RUN CGO_ENABLED=0 go build -o api .

FROM alpine:3.24
COPY --from=builder /api/api .
RUN chmod +x api
EXPOSE 8080
CMD ["./api"]


