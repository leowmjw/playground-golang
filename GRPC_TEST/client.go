package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "github.com/leowmjw/playground-golang/GRPC_TEST/services"
)

func main() {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:3000", opts...)
	if err != nil {
		log.Fatalf("Died in the connectio #%v", err)
	}
	defer conn.Close()
	client := pb.NewMyRouteGuideClient(conn)
	feature, err := client.GetFeature(context.Background(), &pb.Point{Lat: 409146138, Lng: -746188906})
	log.Println(feature.String())
}
