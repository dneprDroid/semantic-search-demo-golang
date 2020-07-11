ARG GO_VERSION=1.11

FROM golang:${GO_VERSION} AS builder

ENV GO111MODULE=on

WORKDIR /main

COPY common/go.* common/
COPY app/go.* app/

WORKDIR /main/app

RUN go mod download

WORKDIR /main

COPY common common/
COPY app/src app/

WORKDIR /main/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

COPY --from=builder /main/app/appserver ./main/

EXPOSE 80
ENTRYPOINT ["/main/appserver"]

