package main

import (
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	pb "github.com/leowmjw/playground-golang/GRPC_TEST/services"
)

type routeGuideServer struct {
	savedFeatures []*pb.Feature
	routeNotes    map[string][]*pb.RouteNote
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}

	// No feature was found, return an unnamed feature
	return &pb.Feature{Location: point}, nil
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.MyRouteGuide_ListFeaturesServer) error {
	return nil
}

func (s *routeGuideServer) RecordRoute(stream pb.MyRouteGuide_RecordRouteServer) error {
	return nil
}

func (s *routeGuideServer) RouteChat(stream pb.MyRouteGuide_RouteChatServer) error {
	return nil
}

func newServer() *routeGuideServer {
	s := new(routeGuideServer)
	s.savedFeatures = make([]*pb.Feature, 0)
	s.routeNotes = make(map[string][]*pb.RouteNote)
	return s
}

func main() {

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen #%v", err)

	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMyRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
