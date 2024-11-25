package main

import (
	as "Replication/AuctionService"
	"context"
	"log"
	"math/rand/v2"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Bidder struct {
	as.UnimplementedAuctionServiceServer
	AuctionContact as.AuctionServiceClient
	MyLatestBid    int64
	Id             int64
	nodes          []string
}

func (bidder *Bidder) Bid(amount int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	result, err := bidder.AuctionContact.Bid(ctx, &as.Amount{
		Amount:   amount,
		Bidderid: bidder.Id,
	})

	if err != nil {
		log.Printf("Failed to bid: %v", err)
		bidder.FindNode()
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Deadline exceeded")
		}
		return
	}
	if !result.Ack {
		bidder.Id = result.Bidderid
		log.Printf("Bid failed")
		return
	}
	bidder.Id = result.Bidderid
	bidder.MyLatestBid = amount
	log.Printf("Bid successful")
}

func (bidder *Bidder) Status() {
	auctionStatus, err := bidder.AuctionContact.Result(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Failed to get auction status: %v", err)
		bidder.FindNode()
		return
		//bidder.Status()
	}

	if auctionStatus.Amount != bidder.MyLatestBid {
		log.Printf("Auction status: %d", auctionStatus.Amount)
		newBid := auctionStatus.Amount + 10
		bidder.Bid(newBid)

	}
}

func (bidder *Bidder) FindNode() {
	for {
		randomNumber := rand.Int64N(2)
		address := "localhost:" + bidder.nodes[randomNumber]
		log.Print(randomNumber)
		log.Print(address)
		conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect to peer %s: %v", address, err)
			continue
		}

		client := as.NewAuctionServiceClient(conn)
		bidder.AuctionContact = client
		log.Printf("Connected to node %s", address)
		return
	}
}

func start() {
	bidder := Bidder{
		nodes:       []string{"5001", "5002", "5003", "5004", "5005"},
		MyLatestBid: -1,
	}

	bidder.FindNode()

	for {
		bidder.Status()
		time.Sleep(4 * time.Second)
	}
}

func main() {
	logFile, err := os.OpenFile("../Server-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	start()
}
