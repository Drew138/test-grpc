package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/drew138/test-grpc/src/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	config := &tls.Config{}
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := messages.NewImageServiceClient(conn)
	user := &messages.User{Id: 1, Provider: "hola", IsDemo: true}
	f, _ := os.Open("./img.jpg")
	defer f.Close()

	fileInfo, _ := f.Stat()
	var size int64 = fileInfo.Size()
	byt := make([]byte, size)
	fmt.Println("hola")
	buffer := bufio.NewReader(f)
	buffer.Read(byt)

	res, err := c.ApplySharpenFilter(context.Background(), &messages.ImageRequest{User: user, Image: byt, Kernel: nil})

	if err != nil {
		fmt.Printf("Error while transforming image %v\n", err)
		return
	}

	fmt.Println("Finished processing image")
	a, _ := os.Create("lena_converted.jpg")
	img, _, _ := image.Decode(bytes.NewReader(res.Image))
	jpeg.Encode(a, img, nil)
}
