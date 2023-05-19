package client

import (
	"book/book_go_author_service/config"
	"book/book_go_author_service/genproto/book_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	BookService() book_service.BookServiceClient
	//AuthorService() author_service.AuthorServiceClient
}

type grpcClients struct {
	bookService book_service.BookServiceClient
	//authorService author_service.AuthorServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connBookService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: book_service.NewBookServiceClient(connBookService),
		// authorService: author_service.NewAuthorServiceClient(connAuthorService),
	}, nil
}

func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}
