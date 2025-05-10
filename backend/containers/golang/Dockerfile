FROM golang:1.23.6

ENV GO_VERSION=1.23.6 \
    GOBIN=/go/bin \
    PATH=$PATH:/go/bin

WORKDIR /usr/local/go/src/golang_clean_architecture/backend

RUN go install github.com/go-delve/delve/cmd/dlv@v1.24.2 && \
    go install github.com/cosmtrek/air@v1.49.0

COPY backend/containers/golang/entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh
RUN go mod download

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
