package main

import (
	"log"

	"cin.ufpe.br/~if998/cittapet/mapservice/proto"
)

func main() {
	log.Println("Hello World!")

	event := &proto.BusEvent{
		Unit:      "108",
		Timestamp: "2000-01-01T12:00:00+00:00",
		Coordinate: &proto.UTM{
			Easting:  293868,
			Northing: 9115080,
		},
	}
	log.Println(event)
}
