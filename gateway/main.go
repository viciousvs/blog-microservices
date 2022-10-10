package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/viciousvs/blog-microservices/gateway/config"
	httpGW "github.com/viciousvs/blog-microservices/gateway/server/http"
	"github.com/viciousvs/blog-microservices/gateway/server/http/routes"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// loads values from gateway.env into the system
	if err := godotenv.Load("gateway.env"); err != nil {
		log.Print("No gateway.env file found")
	}
}
func main() {
	cfg := config.NewConfig()
	//client, err := postClient.NewPostClientGRPC(cfg.PostClientConfig)
	//if err != nil {
	//	log.Fatal(err)
	//}

	conn, err := grpc.Dial(cfg.PostClientConfig.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	client := pbPost.NewPostServiceClient(conn)
	//handler := httpGW.NewHandler(client)
	//r := handler.InitRoutes()

	r := routes.NewMux(client).InitPostRoutes()
	srv := new(httpGW.Server)
	go func() {
		fmt.Println(cfg.ServerConfig.Addr)
		if err := srv.Run(cfg.ServerConfig, r); err != nil {
			log.Fatalf("Cannot run server: %v", err)
		}
	}()

	log.Printf("server started, port:%s", cfg.ServerConfig.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}

}
