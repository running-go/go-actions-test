FROM alpine:3.17 as root-certs
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app


FROM golang:1.20 as builder
WORKDIR /youtube-api-files
COPY --from=root-certs /etc/ssl/certs/ca-certificates /etc/ssl/certs
# 复制源代码到golang镜像中
COPY . .
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64

