

build:
    go build -o ./bin/eolas ./cmd/eolas/

run: build
    ./bin/eolas
