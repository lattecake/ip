package endpoint

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/lattecake/ip/pkg/service"
)

// LocationRequest collects the request parameters for the Location method.
type LocationRequest struct {
	Ip string `json:"ip"`
}

// LocationResponse collects the response parameters for the Location method.
type LocationResponse struct {
	Body service.Body `json:"body"`
	Err  error        `json:"err"`
}

// MakeLocationEndpoint returns an endpoint that invokes Location on the service.
func MakeLocationEndpoint(s service.IpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LocationRequest)
		body, err := s.Location(ctx, req.Ip)
		return LocationResponse{
			Body: body,
			Err:  err,
		}, nil
	}
}

// Failed implements Failer.
func (r LocationResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Location implements Service. Primarily useful in a client.
func (e Endpoints) Location(ctx context.Context, ip string) (body service.Body, err error) {
	request := LocationRequest{Ip: ip}
	response, err := e.LocationEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LocationResponse).Body, response.(LocationResponse).Err
}
