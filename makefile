server:
    go build -o bin/server server/server.go

client:
    go build -o bin/client client/client.go

remote:
    go build -o bin/remote rc/remote.go