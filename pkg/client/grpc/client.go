package grpc

import (
	"google.golang.org/grpc"
)

func NewClientConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("", grpc.WithBalancerName("gorift"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
