package main

import (
	"context"

	"cin.ufpe.br/~if998/cittapet/mapservice/proto"
)

// server implements MapServer.
type mapServer struct{}

// GetBusesLocation is implementing StorageServer.
func (s *mapServer) GetBusesLocation(ctx context.Context, req *proto.GetBusesLocationRequest) (*proto.GetBusesLocationResponse, error) {
	GetBusesLocationResponse := &proto.GetBusesLocationResponse{
		BusesLocation: []*proto.GetBusesLocationResponse_BusLocation{
			&proto.GetBusesLocationResponse_BusLocation{
				Unit:      "108",
				Timestamp: "2000-01-01T12:00:00+00:00",
				Coordinate: &proto.LatLng{
					Latitude:  -7.998191,
					Longitude: -34.87021038376772,
				},
			},
		},
	}
	return GetBusesLocationResponse, nil
}
