package main

import (
	as "Mutual_Exclusion/m/v2/raalgo"
	"context"
	"log"
	"math/rand/v2"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Bidder struct {
	as.UnimplementedAuctionServer
	AuctionContact as.AuctionClient
	MyLatestBid    int64
	Id             int64
	nodes          []string

	peers         map[string]pb.P2PnetworkClient // map of peer addresses to clients
	peerLock      sync.RWMutex
	ME            int64
	N             int64
	Our_Timestamp int64
}

func (bidder *Bidder) Bid(amount int64) {
	bidder.AuctionContact.Bid(context.Background(), &as.Amount{
		Amount: amount,
		Id:     bidder.Id,
	})
}

func (bidder *Bidder) Status() {
	auctionStatus, err := bidder.AuctionContact.Result(context.Background(), &emptypb.Empty{})
	if err != nil {
		FindNode()
		//bidder.Status()
	}
	if auctionStatus.Amount != bidder.MyLatestBid {
		bidder.MyLatestBid = auctionStatus.Amount + 10
		bidder.Bid(bidder.MyLatestBid)
	}
}

func (bidder *Bidder) FindNode() {
	for {
		randomNumber := rand.Int64N(4)
		address := bidder.nodes[randomNumber]
		conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect to peer %s: %v", address, err)
			continue
		}

		client := pb.NewP2PnetworkClient(conn)
		n.peerLock.Lock()
		n.peers[address] = client
		n.peerPorts[nodeid] = address
		n.peerLock.Unlock()

		log.Printf("Connected to node %s", address)

	}
}
func start() {
	bidder := Bidder{
		nodes: []string{"5001", "5002", "5003", "5004", "5005"},
	}
	bidder.FindNode()
	bidder.Status()
}
