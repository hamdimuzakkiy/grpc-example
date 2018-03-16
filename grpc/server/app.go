package main

import (
	"errors"
	"log"
	"net"

	pb "github.com/hamdimuzakkiy/grpc-example/grpc/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l, err := net.Listen("tcp4", ":50053")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer(
		// grpc.UnaryInterceptor(unaryInterceptor),
	)

	if err := new().Register(server); err != nil {
		log.Fatalln(err)
	}

	reflection.Register(server)

	log.Println("Serving ...")

	if err := server.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

type Server struct {
}

func new() Server {
	return Server{}
}

type Calculator struct {
}

func (s Server) Register(server *grpc.Server) error {
	calc := s.NewCalculator()
	pb.RegisterCalculatorServer(server, &calc)
	return nil
}

func (s Server) NewCalculator() Calculator {
	return Calculator{}
}

func (c Calculator) Devide(ctx context.Context, in *pb.DevideRequest) (res *pb.DevideResponse, err error) {
	if in == nil {
		return res, errors.New("nil input")
	}

	if in.Number2 == 0 {
		return res, errors.New("devide by 0")
	}

	return &pb.DevideResponse{Result: in.Number1 / in.Number2}, nil
}

func (c Calculator) Plus(ctx context.Context, in *pb.PlusRequest) (res *pb.PlusResponse, err error) {
	if in == nil {
		return res, errors.New("nil input")
	}

	return &pb.PlusResponse{Result: in.Number1 + in.Number2}, nil
}

func (c Calculator) Minus(ctx context.Context, in *pb.MinusRequest) (res *pb.MinusResponse, err error) {
	if in == nil {
		return res, errors.New("nil input")
	}

	return &pb.MinusResponse{Result: in.Number1 - in.Number2}, nil
}

func (c Calculator) MultiplePlus(ctx context.Context, in *pb.MultiplePlusRequest) (res *pb.PlusResponse, err error) {
	if in == nil {
		return res, errors.New("nil input")
	}

	var result int64
	for _, val := range in.Number {
		result += val
	}

	return &pb.PlusResponse{Result: result}, nil
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := auth(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if len(md["username"]) > 0 && len(md["password"]) > 0 && md["username"][0] == "admin" && md["password"][0] == "admin123"{
			return nil
		}

		return errors.New("authenticate problem")
	}
	return nil
}