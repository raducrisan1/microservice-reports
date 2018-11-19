package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/raducrisan1/microservice-reports/stockreport"
	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(msg)
	}
}

func main() {
	persistAddr := os.Getenv("PERSIST_GRPC_ADDR")
	conn, err := grpc.Dial(persistAddr, grpc.WithInsecure())
	failOnError(err, "Could not connect to microservice-persist gRPC server")
	defer conn.Close()

	//prepares a client to stockpersist microservice
	stockReportClient := stockreport.NewStockReportDataServiceClient(conn)

	impulse := make(chan int, 2)

	go func() {
		impulse <- 1
	}()

	mongoAddr := os.Getenv("MONGO_ADDR")
	mng, err := mongo.NewClient(mongoAddr)
	failOnError(err, "Could not create a mongodb client")
	defer mng.Disconnect(context.Background())

	go func() {
		newGrpcDataServer(mng)
	}()

	err = mng.Connect(context.Background())
	failOnError(err, "Could not connect do mongodb server")

	fmt.Printf("The reports node started\n")
	ticker := time.Tick(time.Second * 30)
	osstop := ossignal()
	stop := false
	coll := mng.Database("tradingdw").Collection("reportdata")

	for !stop {
		select {
		case <-ticker:
			impulse <- 1
		case <-impulse:
			startTime, _ := time.Parse(time.RFC3339, "2018-11-10 09:30Z")
			endTime, _ := time.Parse(time.RFC3339, "2018-11-10 10:00Z")
			req := new(stockreport.StockReportRequest)
			req.Start = startTime.Unix()
			req.End = endTime.Unix()
			req.Resolution = 300

			//every 30 seconds, a call to stockinfo microservice is made.
			res, err := stockReportClient.GetStockReportData(context.Background(), req)
			if err != nil {
				fmt.Printf("An error occurred receiving data from microservice-persist: %s\n", err)
				continue
			}
			records := make([]interface{}, len(res.StockData))
			for index, item := range res.StockData {
				annotatedMsg := *item
				records[index] = annotatedMsg
			}
			_, newErr := coll.InsertMany(context.Background(), records)
			if newErr != nil {
				fmt.Printf("An error occurred saving DW: %s\n", newErr)
			}

		case <-osstop:
			stop = true
			fmt.Println("\nNode stop has been requested")
		}
	}

	fmt.Println("\nThe node has stopped")
}
