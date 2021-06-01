# Stage: build
FROM golang:1.16.4-buster AS buildStage

ENV GOBIN=/go/bin
WORKDIR /go/src

COPY ./ ./

RUN go install -mod=vendor ./...

# Stage: app
FROM golang:1.16.4-buster AS app

ENV GWE_PORT="8080"

COPY --from=buildStage /go/bin/* /usr/local/bin/

WORKDIR /usr/local/bin

CMD ["go_webserver_example"]