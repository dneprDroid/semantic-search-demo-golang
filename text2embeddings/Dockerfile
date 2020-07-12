ARG GO_VERSION=1.18

FROM golang:${GO_VERSION} AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY common/go.* common/
COPY text2embeddings/go.* text2embeddings/

WORKDIR /app/text2embeddings

RUN go mod download

WORKDIR /app

COPY common common/
COPY text2embeddings/*.go text2embeddings/

WORKDIR /app/text2embeddings

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o worker

FROM scratch

ARG TOKENIZER_VERSION=0.1.16

COPY --from=builder /app/text2embeddings/worker /app/text2embeddings/
COPY --from=builder /go/pkg/mod/github.com/sugarme/tokenizer@v${TOKENIZER_VERSION}/pretrained \
                    /go/pkg/mod/github.com/sugarme/tokenizer@v${TOKENIZER_VERSION}/pretrained

ENTRYPOINT ["/app/text2embeddings/worker"]