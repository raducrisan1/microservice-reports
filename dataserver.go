package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/raducrisan1/microservice-api/tradesuggest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	MongoDB *mongo.Client
}

func newGrpcDataServer(mng *mongo.Client) {
	lis, err := net.Listen("tcp", ":3070")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	tradesuggest.RegisterTradeSuggestServiceServer(s, &grpcServer{
		MongoDB: mng})

	//this is used to allow API inspection via grpc_cli tool
	reflection.Register(s)
	stop := ossignal()
	go func() {
		<-stop
		s.Stop()
	}()
	fmt.Println("StockReports node started listening for incoming connections on port 3070")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	fmt.Println("\nThe server has stopped")
}

func (s *grpcServer) GetSuggestions(ctx context.Context, ts *tradesuggest.TradeSuggestRequest) (*tradesuggest.TradeSuggestResponse, error) {

	coll := s.MongoDB.Database("tradingdw").Collection("reportdata")

	var pipe []interface{}
	pipe = append(pipe, bson.M{"$limit": 4})
	cursor, errFind := coll.Aggregate(ctx, pipe)

	if errFind != nil {
		fmt.Printf("Error calling mongo.find: %v", errFind)
	}

	sd := make([]*tradesuggest.Suggestion, 0)

	for cursor.Next(ctx) {
		doc := bsonx.Doc{}

		err := cursor.Decode(&doc)
		if err != nil {
			continue
		}

		item := new(tradesuggest.Suggestion)
		item.Stockname = doc.Lookup("stockname").StringValue()
		item.Rating = doc.Lookup("rating").Int32()
		sd = append(sd, item)
	}

	rsp := tradesuggest.TradeSuggestResponse{
		Suggestions: sd}

	return &rsp, nil

}