package main

import (
	"log"
	"net"

	"cin.ufpe.br/~if998/cittapet/mapservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	log.Println(UTM2LatLng(event.Coordinate))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := &mapServer{}

	c, err := NewConsumer(server)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer c.Close()
	go func(c *consumer) {
		log.Fatal(c.Consume())
	}(c)

	s := grpc.NewServer()
	proto.RegisterMapServer(s, server)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Fatal(s.Serve(lis))
}
