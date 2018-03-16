package main

import (
	"log"
	"time"
	"net/http"
	"bytes"

	"github.com/hamdimuzakkiy/grpc-example/grpc/client"
	pb "github.com/hamdimuzakkiy/grpc-example/grpc/protos"
	"golang.org/x/net/context"
)

const server string = "172.31.4.188"

func main() {
	module, err := client.New(client.Config{
		Username: "admin1",
		Password: "admin123",
		Address:  server+":50053",
	})
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(time.Second)
	tot := 1000

	t := time.Now()
	for i := 0; i < tot; i++ {
		_, _ = module.Calculator.Plus(context.Background(), &pb.PlusRequest{
			Number1: 123,
			Number2: 123,
		})
	}
	grpcRes := time.Now().Sub(t)

	t2 := time.Now()
	clientHttp := &http.Client{}
	for i := 0; i < tot; i++ {
		body := []byte(`{"number1:123, number2:123"}`)
		req, _ := http.NewRequest("POST", "http://"+server+":9010/plus", bytes.NewReader(body))
		_, _ = clientHttp.Do(req)
	}

	resRes := time.Now().Sub(t2)

	log.Println("GRPC : ", grpcRes)
	log.Println("Rest : ", resRes)
}
