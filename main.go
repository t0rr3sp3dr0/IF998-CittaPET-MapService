package main

import (
	"log"
	"net"

	"cin.ufpe.br/~if998/cittapet/mapservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	addr = ":50051"
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

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterMapServer(s, &mapServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Fatal(s.Serve(lis))
}
