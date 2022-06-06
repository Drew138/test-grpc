package main

import (
	"crypto/tls"
	"log"
	"net"

	imageService "github.com/drew138/test-grpc/src/image-service"
	"github.com/drew138/test-grpc/src/messages"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := "50051"

	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Started listening on port: %v\n", port)

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("golang-autocert"),
		HostPolicy: autocert.HostWhitelist("localhost"),
	}
	// go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	creds := credentials.NewTLS(&tls.Config{GetCertificate: certManager.GetCertificate})
	server := grpc.NewServer(grpc.Creds(creds))
	// server := grpc.NewServer()

	reflection.Register(server)

	imageServiceServer := &imageService.ImageServiceServer{}
	messages.RegisterImageServiceServer(server, imageServiceServer)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
