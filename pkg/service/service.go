package service

import (
	"context"
	"github.com/golang/go/src/pkg/encoding/json"
)

// IpService describes the service.
type IpService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Location(ctx context.Context, ip string) (body Body, err error)
}

type Body struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Region  string `json:"region"`
}

type basicIpService struct{}

func (b *basicIpService) Location(ctx context.Context, ip string) (body Body, err error) {
	// TODO implement the business logic of Location
	data := `{"country": "中国", "city": "北京", "region": "朝阳"}`
	err = json.Unmarshal([]byte(data), &body)
	if err != nil {
		return
	}

	return body, err
}

// NewBasicIpService returns a naive, stateless implementation of IpService.
func NewBasicIpService() IpService {
	return &basicIpService{}
}

// New returns a IpService with all of the expected middleware wired in.
func New(middleware []Middleware) IpService {
	var svc IpService = NewBasicIpService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
