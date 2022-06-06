FROM golang:1.17-alpine AS builder

RUN apt-get update

RUN apk add --no-cache ca-certificates git

WORKDIR /graphics-app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

FROM scratch AS final

COPY --from=builder /graphics-app /graphics-app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 50051

VOLUME ["/cert-cache"]

ENTRYPOINT ["/app"]
