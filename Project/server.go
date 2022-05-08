package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/DameInCode/Project/order"
	orderpb "github.com/DameInCode/Project/order/rpc"
	"github.com/DameInCode/Project/book"
	bookpb "github.com/DameInCode/Project/book/rpc"
	"github.com/DameInCode/Project/store"
	"google.golang.org/grpc"
)

func main() {
	log := log.Default()
	ps := store.NewBookStore()
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
	prd := book.NewService(ps)
	bookpb.RegisterServiceServer(srv, prd)

	ord := order.NewService(ords, ps)
	orderpb.RegisterServiceServer(srv, ord)

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
