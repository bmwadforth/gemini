package main

import (
	"flag"
	"fmt"
	GeminiService "gemini/gemini_service"
	"gemini/middleware"
	GeminiServicePB "gemini/protocol_buffers/gemini_service"
	"gemini/util"
	"google.golang.org/grpc"
	"net"
	"os"
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
		util.SLogger.Panicf("Failed to listen on port:%d : %v", *port, err)
	}

	util.SLogger.Infof("Listening on port :%d", *port)
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(middleware.ApiKeyAuthenticationInterceptor),
	}
	s := grpc.NewServer(opts...)
	GeminiServicePB.RegisterGeminiServer(s, &GeminiService.Server{})
	if err := s.Serve(lis); err != nil {
		util.SLogger.Panicf("Failed to start server: %v", err)
	}
}
