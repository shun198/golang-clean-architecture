FROM --platform=linux/x86_64 golang:1.23

WORKDIR /usr/local/go/src/golang_clean_architecture/backend

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@v1.42.0

COPY backend/containers/golang/entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh
RUN go mod download

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
