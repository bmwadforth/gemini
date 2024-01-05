package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	BookService "svc-template/book_service"
	BookServicePB "svc-template/protocol_buffers/book_service"
	"svc-template/util"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	logCleanup := util.InitLogger()
	defer logCleanup(util.Logger)
	flag.Parse()

	util.IsProduction = os.Getenv("APP_ENV") == "PRODUCTION"
	if util.IsProduction {
		util.LoadEnvironmentVariables()
	} else {
		util.LoadConfiguration()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		util.SLogger.Errorf("Failed to listen on port:%d : %v", *port, err)
	}

	util.SLogger.Infof("Listening on port :%d", *port)

	s := grpc.NewServer()
	BookServicePB.RegisterBookServiceServer(s, &BookService.Server{})
	if err := s.Serve(lis); err != nil {
		util.SLogger.Errorf("Failed to start server: %v", err)
	}
}
