FROM golang:1.15.6-alpine@sha256:14d4ada6aed6ed019a2b609108d8b02ca09d6c338f6499397c52148d22230967
LABEL maintainer=eltonribeiro@outlook.com
LABEL app=backend
EXPOSE 8080
ENV CGO_ENABLED=0 \
    GOOS="linux" \
    GOARCH="amd64" \
    GO111MODULE=on

WORKDIR /app
COPY . .
RUN apk --no-cache add curl
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
ENTRYPOINT ["air"]