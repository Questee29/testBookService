package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Questee29/testBookService/configs"
	"github.com/Questee29/testBookService/database"
	"github.com/Questee29/testBookService/middleware"
	grpcServer "github.com/Questee29/testBookService/pkg/grpc"
	grpcHandlers "github.com/Questee29/testBookService/pkg/grpc/handler"
	handlers "github.com/Questee29/testBookService/pkg/handlers"
	Repository "github.com/Questee29/testBookService/pkg/repository"
	Servc "github.com/Questee29/testBookService/pkg/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//init env config
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	log.Println(config.Database.Host)
	//init database
	db, err := database.New()
	if err != nil {
		log.Println(err)
		log.Fatalln(errors.New(`failed to load database`))
	}

	//Init repository to work with db
	Repository := Repository.NewRepository(db)

	//Init REST service to work with logic
	Service := Servc.NewRestService(Repository)
	//init GRPC service to work with logic
	ServiceGrpc := Servc.NewGrpcSerivce(Repository)

	//Init REST handlers
	handlerGetBook := handlers.NewHandler(Service)
	handlerGetAuthor := handlers.NewGetAuthorHandler(Service)

	//Init GRPC handlers and Init GRPC server
	handlerGRPC := grpcHandlers.NewGrpcHandler(ServiceGrpc)
	grpcServ := grpcServer.NewServer(grpcServer.Deps{
		GrpcHandler: handlerGRPC,
	})

	//Handler rest functions
	http.Handle("/get-book", middleware.SetContentTypeJSON(handlerGetBook))
	http.Handle("/get-author", middleware.SetContentTypeJSON(handlerGetAuthor))

	//go REST server
	go func() {
		log.Println("Starting listening REST server")
		http.ListenAndServe(config.Server.Port, nil)
	}()

	//go GRPC server
	go func() {
		log.Println("Starting listening GRPC server")
		if err := grpcServ.ListenAndServe(config.Server.GrpcPort); err != nil {
			log.Printf("grpc ListenAndServe err %s", err)
		}
	}()

	//waiting for exit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("shutting down server")
	db.Close()

}
