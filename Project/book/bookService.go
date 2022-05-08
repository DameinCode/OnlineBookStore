package book

import (
	"context"

	bookrpc "github.com/DameInCode/Project/book/rpc"

	"github.com/DameInCode/Project/store"
)

type service struct {
	bookrpc.UnimplementedServiceServer
	s store.BookStore
}

func NewService(s store.BookStore) *service {
	return &service{s: s}
}

func (s *service) ListBooks(ctx context.Context, r *bookrpc.ListBooksRequest) (*bookrpc.ListBooksResponse, error) {
	return &bookrpc.ListBooksResponse{Books: s.s.Books()}, nil
}
func (s *service) BookOfId(ctx context.Context, r *bookrpc.BookOfIdRequest) (*bookrpc.BookOfIdResponse, error) {
	prod, err := s.s.Book(r.Id)
	return &bookrpc.BookOfIdResponse{Book: prod}, err
}
