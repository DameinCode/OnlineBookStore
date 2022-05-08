package order

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"context"

	orderrpc "github.com/DameInCode/Project/order/rpc"
	"github.com/DameInCode/Project/store"
)

type service struct {
	orderrpc.UnimplementedServiceServer
	s  store.OrderStore
	s2 store.BookStore
}

func NewService(s store.OrderStore, s2 store.BookStore) *service {
	return &service{s: s, s2: s2}
}

func (s *service) ListOrders(ctx context.Context, r *orderrpc.ListOrdersRequest) (*orderrpc.ListOrdersResponse, error) {
	return &orderrpc.ListOrdersResponse{Orders: s.s.Orders()}, nil
}
type Properties struct {
	Name     string
	Postcode string
	City     string
	Score    float64
}

type Features struct {
	Properties Properties
}

type TargetResponse struct {
	Features []Features
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func gen_id(lName string) (id string) {
	id = strconv.Itoa(rand.Intn(9000)+1000) + "-" + time.Now().Format("20060102-150405") + "-" + strings.ToUpper(lName)
	return
}

const APIADDPREFIX string = "https://api-adresse.data.gouv.fr/search/?q="
const APIADDSUFIX string = "&limit=1"
const MIN_SCORE float32 = 0.3

func (s *service) CreateOrder(ctx context.Context, r *orderrpc.CreateOrderRequest) (*orderrpc.CreateOrderResponse, error) {
	for key := range r.BookList {
		_, err := s.s2.Book(key)
		if err != nil {
			return &orderrpc.CreateOrderResponse{Order: nil}, errors.New("book " + key + " not available in store, order was not created")
		}
	}
	resp := TargetResponse{} //expected body answer from the api
	linename := strings.Replace(r.Address.Street, " ", "+", -1)
	city := strings.Replace(r.Address.City, " ", "+", -1)
	txt_request := linename + ",+" + r.Address.Postal + ",+" + city
	request := APIADDPREFIX + txt_request + APIADDSUFIX
	getJson(request, &resp) 
	if len(resp.Features) == 0 || resp.Features[0].Properties.Score < float64(MIN_SCORE) {
		return &orderrpc.CreateOrderResponse{Order: nil}, errors.New("address not found")
	}
	address := orderrpc.Address{Street: resp.Features[0].Properties.Name, Postal: resp.Features[0].Properties.Postcode, City: resp.Features[0].Properties.City, Country: "France"}
	gid := gen_id(r.LName)
	ord := orderrpc.Order{
		Id:          gid,
		FName:       r.FName,
		LName:       r.LName,
		BookList: r.BookList,
		Address:     &address,
	}
	s.s.SetOrder(&ord) 
	return &orderrpc.CreateOrderResponse{Order: &ord}, nil
}
