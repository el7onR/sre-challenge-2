FROM golang:1.15.6-buster@sha256:27a4178f1a02c2a2ca0070bd4be92b95fd9cbdd76eb49b54a688c08003bcb26f AS builder
LABEL maintainer=eltonribeiro@outlook.com
LABEL app=backend
EXPOSE 8080
ENV CGO_ENABLED=1 \
    GOOS="linux" \
    GOARCH="amd64" \
    GO111MODULE=on

WORKDIR /build
COPY . .
RUN go build -o backend cmd/main.go

FROM golang:1.15.6-buster@sha256:27a4178f1a02c2a2ca0070bd4be92b95fd9cbdd76eb49b54a688c08003bcb26f
WORKDIR /app
COPY --from=builder /build/backend /app/backend
ENTRYPOINT ["./backend"]