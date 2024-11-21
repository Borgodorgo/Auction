package main

import (
	as "Replication/m/v2/AuctionService/Auction"
	"context"
	"log"
	"math/rand/v2"

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
	log.Printf("Bidding %d", amount)
	bidder.AuctionContact.Bid(context.Background(), &as.Amount{
		Amount:   amount,
		Bidderid: bidder.Id,
	})
}

func (bidder *Bidder) Status() {
	auctionStatus, err := bidder.AuctionContact.Result(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Failed to get auction status: %v", err)
		bidder.FindNode()
		//bidder.Status()
	}

	if auctionStatus.Amount != bidder.MyLatestBid {
		log.Printf("Auction status: %d", auctionStatus.Amount)
		bidder.MyLatestBid = auctionStatus.Amount + 10
		bidder.Bid(bidder.MyLatestBid)
	}
}

func (bidder *Bidder) FindNode() {
	for {
		randomNumber := rand.Int64N(4)
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
	log.Print("Found node")
	for {
		bidder.Status()
	}
}
