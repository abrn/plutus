# golang:alpine
FROM golang@sha256:3e935ab77ba5d71c7778054fbb60c029c1564b75266beeeb4223aa04265e16c1 AS builder

# install git
RUN apk update && apk add --no-cache git

ENV USER=appuser
ENV UID=10001
# set up the linux user
RUN adduser \
        --disabled-password \
        --gecos "" \
        --home "/nonexistent" \
        --shell "/sbin/nologin" \
        --no-create-home \
        --uid "${UID}" \
        "${USER}"
# set environment variables for golang
ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

WORKDIR /build

# copy over code and build the binary
COPY ./code .
RUN go mod download
RUN go mod verify
RUN go build -ldflags="-w -s" -o ./binary

# switch to a super lightweight linux container
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# copy over the server binary only
COPY --from=builder /build/binary binary

USER appuser:appuser

# change this port to
EXPOSE 8443

ENTRYPOINT ["./binary"]
