FROM golang:1.21.6 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-s -w -extldflags "-static"' -o goapp
############################################
FROM alpine:3.19

RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*
ENV TZ=Asia/Bangkok

WORKDIR /app

COPY ./configs ./configs
COPY --from=builder /app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]
