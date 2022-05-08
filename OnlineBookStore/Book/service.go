package book

import (
	"context"

	bookrpc "github.com/DameinCode/OnlineBookStore/OnlineBookStore/Book/rpc"

	"github.com/DameinCode/OnlineBookStore/OnlineBookStore/store"
)

// Service holds RPC handlers for the product service. It implements the product.ServiceServer interface.
type service struct {
	bookrpc.UnimplementedServiceServer
	s store.BookStore
}

func NewService(s store.BookStore) *service {
	return &service{s: s}
}

// Fetch all existing products in the system.
func (s *service) ListBooks(ctx context.Context, r *bookrpc.ListBooksRequest) (*bookrpc.ListBooksResponse, error) {
	return &bookrpc.ListBooksResponse{Books: s.s.Books()}, nil
}
func (s *service) BookOfId(ctx context.Context, r *bookrpc.BookOfIdRequest) (*bookrpc.BookOfIdResponse, error) {
	boo, err := s.s.Book(r.Id)
	return &bookrpc.BookOfIdResponse{Book: boo}, err
}
