FROM golang:1.15.6-alpine@sha256:14d4ada6aed6ed019a2b609108d8b02ca09d6c338f6499397c52148d22230967 AS builder
LABEL maintainer=eltonribeiro@outlook.com
LABEL app=backend
EXPOSE 8080
ENV CGO_ENABLED=0 \
    GOOS="linux" \
    GOARCH="amd64" \
    GO111MODULE=on

WORKDIR /build
COPY . .
RUN go build -o backend cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /build/backend /app/backend
ENTRYPOINT ["./backend"]