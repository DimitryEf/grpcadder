// Golang and gRPC
//
//скачиваем https://github.com/protocolbuffers/protobuf/releases/tag/v3.10.0
//protoc и добавляем в path переменныу окружения путь до bin
//
//создаем adder.proto (например)
//в консоли набираем:
//protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
//
//создаем grpcserver.go и main.go
//
//для теста скачиваем https://github.com/ktr0731/evans/releases
//в новой консоли выбираем C:\tools\evans\evans.exe api/proto/adder.proto -p 8080
//
//запускаем клиента
//go run ./cmd/client/main.go 1 2
//

package adder

import (
	"context"
	"github.com/DimitryEf/grpcadder/pkg/api"
)

type GRPCServer struct {

}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error){
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}