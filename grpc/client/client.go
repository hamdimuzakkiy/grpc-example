package client

import (
	"time"

	pb "github.com/hamdimuzakkiy/grpc-example/grpc/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
)

type Client struct {
	Calculator pb.CalculatorClient
}

type Config struct {
	Username string
	Password string
	Address  string
}

type loginCreds struct {
	Username, Password string
}

func (c *loginCreds) RequireTransportSecurity() bool {
	return false
}

func (c *loginCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}, nil
}

func New(c Config) (Client, error) {
	r, err := naming.NewDNSResolverWithFreq(time.Second * 1)
	if err != nil {
		return Client{}, err
	}

	conn, err := grpc.Dial(c.Address, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&loginCreds{
		Username: c.Username,
		Password: c.Password,
	}), grpc.WithBalancer(grpc.RoundRobin(r)))
	if err != nil {
		return Client{}, err
	}

	return Client{
		Calculator: pb.NewCalculatorClient(conn),
	}, nil
}
