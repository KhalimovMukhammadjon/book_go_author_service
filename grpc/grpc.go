package grpc

import (
	"book/book_go_author_service/config"
	"book/book_go_author_service/genproto/author_service"
	"book/book_go_author_service/grpc/client"
	"book/book_go_author_service/grpc/service"

	"book/book_go_author_service/pkg/logger"
	"book/book_go_author_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	author_service.RegisterAuthorServiceServer(grpcServer, service.NewAuthorService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
