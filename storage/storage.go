package storage

import (
	"book/book_go_author_service/genproto/author_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Author() AuthorRepoI
}

type AuthorRepoI interface {
	Create(ctx context.Context, req *author_service.CreateAuthorRequest) (resp *author_service.AuthorPrimaryKey, err error)
	GetById(ctx context.Context, req *author_service.GetAuthorListRequest) (resp *author_service.Author, err error)
	GetList(ctx context.Context, req *author_service.GetAuthorListRequest) (resp *author_service.GetAuthorListResponse, err error)
	Update(ctx context.Context, req *author_service.UpdateAuthorRequest) (rowsAffected int64, err error)
	PatchUpdate(ctx context.Context, req *author_service.PatchUpdateRequest) (rowsAffected int64, err error)
}
