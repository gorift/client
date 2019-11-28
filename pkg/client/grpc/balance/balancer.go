package balance

import (
	"context"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
)

func init() {
	balancer.Register(newGoriftBuilder())
}

func newGoriftBuilder() balancer.Builder {
	return base.NewBalancerBuilder("gorift", &goriftPickerBuilder{})
}

type goriftPickerBuilder struct{}

func (g *goriftPickerBuilder) Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker {
	return nil
}

type goriftPicker struct{}

func (g *goriftPicker) Pick(ctx context.Context, opts balancer.PickOptions) (conn balancer.SubConn, done func(balancer.DoneInfo), err error) {
	return nil, nil, nil
}
