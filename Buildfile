build: |
  GO111MODULE=on \
  go build -mod=mod -o bin/application ./cmd/web
