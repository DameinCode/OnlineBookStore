package account

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"context"

	accountrpc "github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/Account/rpc"
	"github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/store"
)

// Service holds RPC handlers for the order service. It implements the orderrpc.ServiceServer interface.
type service struct {
	accountrpc.UnimplementedServiceServer
	s  store.AccountStore
	s2 store.BookStore
}

func NewService(s store.AccountStore, s2 store.BookStore) *service {
	return &service{s: s, s2: s2}
}

func (s *service) PersonalAccount(ctx context.Context, p *accountrpc.PersonalAccountRequest) (*accountrpc.PersonalAccountResponse, error) {
	
} 

// Fetch all existing orders in the system.
func (s *service) ListOrders(ctx context.Context, r *accountrpc.ListOrdersRequest) (*accountrpc.ListOrdersResponse, error) {
	return &accountrpc.ListOrdersResponse{Orders: s.s.Orders()}, nil
}

type Carting struct {
	bookList map[string]int
}
//fields of importance and their sub groups division of the address api response
type Properties struct {
	id string
    typ string // admin/user/author
    fname string 
    lname string 
    username string
    pswd string // hash of the password to check 
    cart Carting
}
