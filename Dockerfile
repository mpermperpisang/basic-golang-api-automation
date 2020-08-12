FROM golang:latest
WORKDIR /go/src/github.com/basic-golang-api-automation-docker
RUN GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.10.0 && \
    go get github.com/joho/godotenv && \
    go get github.com/yalp/jsonpath
