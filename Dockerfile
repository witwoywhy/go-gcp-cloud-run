FROM golang:1.21.6 as builder

RUN apk add --no-cache tzdata
ENV TZ=Asia/Bangkok

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goapp
    
############################################
FROM alpine:latest as deploy

WORKDIR /app

COPY ./configs ./configs
COPY --from=builder /app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]