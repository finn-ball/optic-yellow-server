# Optic Yellow

Automate booking tennis courts.
```sh
go run main.go
```

## Protobufs

To generate, run:
```sh
protoc -I="$PWD"/pkg/proto \
--go_opt=paths=source_relative --go_out="$PWD"/pkg/proto/ \
--go-grpc_opt=paths=source_relative --go-grpc_out="$PWD"/pkg/proto/ \
"$PWD"/pkg/proto/optic_yellow.proto
```
