package main

import (
	"Mutual_Exclusion/m/v2/raalgo"
	pb "Mutual_Exclusion/m/v2/raalgo"
	"context"
	"log"
	"sync"
)

type P2PNode struct {
	pb.UnimplementedAuctionServer
	peers             map[string]pb.AuctionClient // map of peer addresses to clients
	peerLock          sync.RWMutex
	leader            pb.AuctionClient
	Highest_Bid       int64
	Highest_BidId     int64
	Our_Timestamp     int64
	Highest_Timestamp int64
	IsLeader          bool
	peerPorts         []string
}

func (n *P2PNode) Bid(bid pb.Amount) *raalgo.Ack {
	if !n.IsLeader {
		ack, _ := n.leader.Bid(bid)
		return ack
	}
	//if (!leader) {return response from leader}
	if bid.Amount > n.Highest_Bid {
		UpdateFollowers(bid)
		n.Highest_Bid = bid.Amount
		n.Highest_BidId = bid.Id
		n.Highest_Timestamp = n.Highest_Timestamp + 1
		return &pb.Ack{
			ack: true,
			id:  n.Highest_BidId,
		}
	}
}

func result() (outcome int) {
	//return highest value
}

func (n *P2PNode) UpdateFollowers(update pb.Amount) {
	n.peerLock.RLock()
	for address := range n.peers {
		if peer, exists := n.peers[address]; exists {
			_, err := peer.Update(context.Background(), update)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
		//Send message to followers with the updated value

	}
	n.peerLock.RUnlock()
}

func (n *P2PNode) Update(ctx context.Context, update pb.Amount) {
	//Call from leader to update value
}

func Election() {
	//Election code
}
