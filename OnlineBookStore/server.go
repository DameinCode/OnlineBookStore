package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/Account"
	accountpb "github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/Account/rpc"
	"github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/Book"
	bookpb "github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/Book/rpc"
	"github.com/DameinCode/Golang-spring-2022/Uni/Final/OnlineBookStore/store"
	"google.golang.org/grpc"
)

func main() {
	log := log.Default()
	ps := store.NewProductStore()
	ords := store.NewOrderStore()

	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	defer srv.Stop()

	//declaration of services
	prd := Book.NewService(ps)
	bookpb.RegisterServiceServer(srv, prd)

	ord := Account.NewService(ords, ps)
	accountpb.RegisterServiceServer(srv, ord)

	go func() {
		log.Println("server started on port 8000...")
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("stopping the server")
}
