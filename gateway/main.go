package main

import (
	"context"
	httpGW "github.com/viciousvs/blog-microservices/gateway/server/http"
	postClient "github.com/viciousvs/blog-microservices/gateway/service/grpc/post"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const PORT = "8081"

func main() {
	conn, err := postClient.Connect("localhost:50051")
	if err != nil {
		log.Fatalf("cannot dial to grpc server: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	client := postClient.GetClient(conn)
	handler := httpGW.NewHandler(client)
	r := handler.InitRoutes()

	srv := new(httpGW.Server)
	go func() {
		if err := srv.Run(PORT, r); err != nil {
			log.Fatalf("Cannot run server: %v", err)
		}
	}()

	log.Printf("server started, port:%s", PORT)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}

}
