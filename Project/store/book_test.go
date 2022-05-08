package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type bookStoreTestSuite struct {
	suite.Suite
	s BookStore
}

func (s *bookStoreTestSuite) SetupTest() {
	s.s = NewBookStore()
}

func TestProductStoreTestSuite(t *testing.T) {
	suite.Run(t, new(bookStoreTestSuite))
}

func (s *bookStoreTestSuite) TestProducts_OK() {
	s.Len(s.s.Books(), 5)
}

func (s *bookStoreTestSuite) TestProduct_OK() {
	pdt, err := s.s.Book("First-book")
	s.Require().NoError(err)
	s.Equal("First-book", pdt.GetId())
}

func (s *bookStoreTestSuite) TestProduct_Err() {
	unknownID := "0001"
	pdt, err := s.s.Book(unknownID)
	s.Nil(pdt)
	s.Error(err)
	s.EqualError(err, "Sorry, book not found")
}
