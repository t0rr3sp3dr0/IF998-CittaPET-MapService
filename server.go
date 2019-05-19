package main

import (
	"context"
	"sync"

	"cin.ufpe.br/~if998/cittapet/mapservice/proto"
)

// server implements MapServer.
type mapServer struct {
	Buses sync.Map
}

// GetBusesLocation is implementing MapServer.
func (s *mapServer) GetBusesLocation(ctx context.Context, req *proto.GetBusesLocationRequest) (*proto.GetBusesLocationResponse, error) {
	response := &proto.GetBusesLocationResponse{
		BusesLocation: []*proto.GetBusesLocationResponse_BusLocation{},
	}

	s.Buses.Range(func(keu interface{}, val interface{}) bool {
		v := val.(*proto.GetBusesLocationResponse_BusLocation)
		response.BusesLocation = append(response.BusesLocation, v)

		return true
	})

	return response, nil
}
