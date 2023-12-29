FROM alpine:3.17 as root-certs
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app


FROM golang:1.20 as builder
WORKDIR /youtube-api-files
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# 复制源代码到golang镜像中
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./youtube-stats ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./youtube-stats ./...


FROM scratch as final
WORKDIR /home/app
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder /youtube-api-files/youtube-stats /youtube-stats
USER app
ENTRYPOINT ["/youtube-stats"]
# CMD ["sleep", "6000"]
