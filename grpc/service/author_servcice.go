package service

import (
	"book/book_go_author_service/config"
	"book/book_go_author_service/genproto/author_service"
	"book/book_go_author_service/grpc/client"
	"book/book_go_author_service/pkg/logger"
	"book/book_go_author_service/storage"
)

type authorService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	author_service.UnimplementedAuthorServiceServer
}

func NewAuthorService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *authorService {
	return &authorService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}
