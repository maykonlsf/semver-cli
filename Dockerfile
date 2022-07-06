FROM golang:1.17-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o semver cmd/semver/main.go

FROM alpine:3
COPY --from=builder /build/semver /bin/
COPY .bashrc /root/.bashrc
RUN apk update && apk add --no-cache git bash openssh curl
CMD ["bash"]
