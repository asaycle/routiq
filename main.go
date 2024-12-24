package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/rand/v2"
	"time"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewTagServiceClient(conn)

	// CreateTag メソッドを呼び出す
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.CreateTag(ctx, &pb.CreateTagRequest{
		Tag: &pb.Tag{
			Name: fmt.Sprintf("テストタグ-%d", rand.IntN(10000)),
		},
	})
	if err != nil {
		log.Fatalf("Error CreateTag: %v", err)
	}

	req := &pb.ListTagsRequest{
		PageSize: 10,
	}
	res, err := client.ListTags(ctx, req)
	if err != nil {
		log.Fatalf("Error calling ListTags: %v", err)
	}

	log.Printf("Response from server: %v", res)

	routeClient := pb.NewRouteServiceClient(conn)
	route, err := routeClient.CreateRoute(ctx, &pb.CreateRouteRequest{
		Route: &pb.Route{
			DisplayName: fmt.Sprintf("Biwaichi-%d", rand.IntN(1000)),
			Description: "初心者の登竜門",
			Feature:     biwaIchiRoute,
		},
	})
	if err != nil {
		log.Fatalf("Error CreateRoute: %v", err)
	}
	log.Printf("CreateRouteResponse: %v", route)

	ridingClient := pb.NewRidingServiceClient(conn)
	resp, err := routeClient.ListRoutes(ctx, &pb.ListRoutesRequest{})
	if err != nil {
		log.Fatalf("Error ListRoutes: %v", err)
	}
	for _, route := range resp.Routes {
		log.Printf("Route: %v", route)
		riding, err := ridingClient.CreateRiding(ctx, &pb.CreateRidingRequest{
			Riding: &pb.Riding{
				Title:   "test",
				RouteId: route.Id,
				Note:    "楽しかったです",
				Score:   0,
				Date:    &date.Date{Year: 2024, Month: 12, Day: 1},
				UserId:  "admin",
			},
		})
		if err != nil {
			slog.Error("failed create riding", slog.Any("err", err))
			continue
		}
		slog.Info("created riding", slog.Any("riding", riding))
	}
}

var biwaIchiRoute = `
{
	"type": "Feature",
	"geometry": {
		"type": "LineString",
		"coordinates": [
			[135.8519112, 34.9928373],
			[135.905369, 34.977973],
			[135.909926, 35.009944],
			[135.946325, 35.116805],
			[136.068269, 35.184707],
			[136.261886, 35.378466],
			[136.206756, 35.500407],
			[136.123444, 35.487701],
			[136.037262, 35.376511],
			[136.016251, 35.297281],
			[135.864812, 35.003792]
		],
		"properties": {
			"name": "Biwaichi Route",
			"description": "A scenic cycling route around Lake Biwa."
		}
	}
}
`
