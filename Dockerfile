FROM golang:alpine AS builder
COPY . /app/
WORKDIR /app
RUN go build -o app


FROM alpine
COPY --from=builder /app/app /app/app
COPY --from=builder /app/templates /app/templates
WORKDIR /app/
RUN apk add curl
HEALTHCHECK CMD curl --fail http://localhost:8080/health || exit 1
CMD ["./app"]