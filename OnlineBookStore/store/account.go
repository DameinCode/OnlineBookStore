package store 

import (
	"sort"
	"sync"

	// "C:/Users/user/Documents/Coding/Golang/Uni/Final/OnlineBookStore/Account/rpc"
	"github.com/DameinCode/OnlineBookStore/OnlineBookStore/Account/rpc"
	"github.com/pkg/errors"
)

// Storage interface for orders.
type AccountStore interface {
	Account(id string) (*accountrpc.Account, error)
	// Retrieve all orders in the store, sorted by id asc.
	Accounts() []*accountrpc.Account
	// Upsert an order in the store.
	SetAccount(*accountrpc.Account)
}

type accountStore struct {
	lock sync.RWMutex
	m    map[string]*accountrpc.Account
}

// Create a new account store.
func NewAccountStore() AccountStore {
	return &accountStore{
		m: make(map[string]*accountrpc.Account),
	}
}

// Retrieve an order by ID from the store.
func (s *accountStore) Orders() []*accountrpc.Account {
	s.lock.RLock()
	defer s.lock.RUnlock()

	ods := make([]*accountrpc.Account, 0, len(s.m))
	for _, o := range s.m {
		ods = append(ods, o)
	}

	sort.Slice(ods, func(i, j int) bool {
		return ods[i].GetId() < ods[j].GetId()
	})

	return ods
}
