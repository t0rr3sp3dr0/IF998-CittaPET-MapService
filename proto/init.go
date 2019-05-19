//go:generate protoc -I . --go_out=plugins=grpc:. event.proto map.proto

package proto
