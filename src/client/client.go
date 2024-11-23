package main

import (
	pb "assignment5/auction/src/grpc"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	serverAddr string
	client     pb.AuctionServiceClient
}

func NewClient(serverAddr string) (*Client, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server %s: %v", serverAddr, err)
	}

	client := pb.NewAuctionServiceClient(conn)
	return &Client{serverAddr: serverAddr, client: client}, nil
}

func (c *Client) PlaceBid(bidder string, amount int32) {
	req := &pb.BidRequest{
		Bidder: bidder,
		Amount: amount,
	}

	resp, err := c.client.Bid(context.Background(), req)
	if err != nil {
		log.Printf("Failed to place bid: %v", err)
		return
	}

	fmt.Printf("Response: %s\n", resp.Message)
}

func (c *Client) GetStatus() {
	req := &pb.StatusRequest{}
	resp, err := c.client.GetStatus(context.Background(), req)
	if err != nil {
		log.Printf("Failed to get status: %v", err)
		return
	}

	fmt.Printf("Highest Bidder: %s\n", resp.HighestBidder)
	fmt.Printf("Highest Bid: %d\n", resp.HighestBid)
	fmt.Printf("Message: %s\n", resp.Message)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run client.go <server_address>")
	}

	serverAddr := os.Args[1]

	client, err := NewClient(serverAddr)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("Auction Client")
	fmt.Println("1: Place a bid")
	fmt.Println("2: Get auction status")
	fmt.Println("3: Exit")

	for {
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			log.Printf("Invalid input: %v", err)
			continue
		}

		switch choice {
		case 1:
			var bidder string
			var amount int32
			fmt.Print("Enter bidder name: ")
			fmt.Scanf("%s", &bidder)
			fmt.Print("Enter bid amount: ")
			fmt.Scanf("%d", &amount)

			client.PlaceBid(bidder, amount)

		case 2:
			client.GetStatus()

		case 3:
			fmt.Println("Exiting client.")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
