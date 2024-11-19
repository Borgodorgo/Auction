package main

import (
	pb "Mutual_Exclusion/m/v2/raalgo"
	sv "Mutual_Exclusion/m/v2/serverside"
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type P2PNode struct {
	pb.UnimplementedAuctionServer
	sv.UnimplementedP2PNetworkServer
	peers             map[string]sv.P2PNetworkClient // map of peer addresses to clients
	peerLock          sync.RWMutex
	leader            sv.P2PNetworkClient
	Highest_Bid       int64
	Highest_BidId     int64
	Our_Timestamp     int64
	Highest_Timestamp int64
	IsLeader          bool
	peerPorts         []string
}

func (n *P2PNode) Bid(ctx context.Context, bid *pb.Amount, err error) (ack *pb.Ack) {
	if !n.IsLeader {
		ack, _ := n.leader.Bid(ctx, &sv.Amount{
			Amount: bid.Amount,
			Id:     bid.Id,
		})

		return &pb.Ack{
			Ack: ack.Ack,
			Id:  ack.Id,
		}
	}
	//if (!leader) {return response from leader}
	if bid.Amount > n.Highest_Bid {
		n.UpdateFollowers(bid)
		n.Highest_Bid = bid.Amount
		n.Highest_BidId = bid.Id
		n.Highest_Timestamp = n.Highest_Timestamp + 1
		return &pb.Ack{
			Ack: true,
			Id:  n.Highest_BidId,
		}
	}
	return &pb.Ack{
		Ack: false,
		Id:  n.Highest_BidId,
	}
}

func (n *P2PNode) result() (outcome int64) {
	//return highest value
	return n.Highest_Bid
}

func (n *P2PNode) UpdateFollowers(update *pb.Amount) {
	n.peerLock.RLock()
	for address := range n.peers {
		if peer, exists := n.peers[address]; exists {
			_, err := peer.Update(context.Background(), &sv.Amount{
				Amount:    update.Amount,
				Id:        update.Id,
				Timestamp: n.Highest_Timestamp})
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
		//Send message to followers with the updated value

	}
	n.peerLock.RUnlock()
}

func (n *P2PNode) Update(ctx context.Context, update *sv.Amount) (ack *sv.Response) {
	//Call from leader to update value
	n.Highest_Bid = update.Amount
	n.Highest_BidId = update.Id
	n.Highest_Timestamp = update.Timestamp
	return &sv.Response{
		Ack: true,
	}
}

func Election() {
	//Election code
}
