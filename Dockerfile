ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder
RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /website-api

FROM alpine:latest AS runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /website-api /website-api
COPY .env ./
EXPOSE 8000
ENTRYPOINT [ "/website-api" ]