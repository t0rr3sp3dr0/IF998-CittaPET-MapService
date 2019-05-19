package main

import (
	"errors"

	"cin.ufpe.br/~if998/cittapet/mapservice/proto"
	"github.com/im7mortal/UTM"
)

const (
	zoneNumber = 25
	zoneLetter = "L"
)

func UTM2LatLng(utm *proto.UTM) (*proto.LatLng, error) {
	if utm == nil {
		return nil, errors.New("utm == nil")
	}

	if utm.Easting == 0 || utm.Northing == 0 {
		return nil, nil
	}

	lat, lng, err := UTM.ToLatLon(utm.Easting, utm.Northing, zoneNumber, zoneLetter)
	if err != nil {
		return nil, err
	}

	return &proto.LatLng{
		Latitude:  lat,
		Longitude: lng,
	}, nil
}
