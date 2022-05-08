package store

import (
	"sort"
	"sync"

	"github.com/DameInCode/Project/order/rpc"
	"github.com/pkg/errors"
)

type OrderStore interface {
	Order(id string) (*orderrpc.Order, error)
	Orders() []*orderrpc.Order
	SetOrder(*orderrpc.Order)
}

type orderStore struct {
	lock sync.RWMutex
	m    map[string]*orderrpc.Order
}

func NewOrderStore() OrderStore {
	return &orderStore{
		m: make(map[string]*orderrpc.Order),
	}
}

func (s *orderStore) Orders() []*orderrpc.Order {
	s.lock.RLock()
	defer s.lock.RUnlock()

	ods := make([]*orderrpc.Order, 0, len(s.m))
	for _, o := range s.m {
		ods = append(ods, o)
	}

	sort.Slice(ods, func(i, j int) bool {
		return ods[i].GetId() < ods[j].GetId()
	})

	return ods
}

func (s *orderStore) Order(id string) (*orderrpc.Order, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	o, ok := s.m[id]
	if !ok {
		return nil, errors.New("order not found")
	}

	return o, nil
}

func (s *orderStore) SetOrder(o *orderrpc.Order) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.m[o.GetId()] = o
}
