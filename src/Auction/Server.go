package main

import (
	as "Replication/m/v2/AuctionService/Auction"
	rs "Replication/m/v2/ReplicationService/Replication"
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type P2PNode struct {
	as.UnimplementedAuctionServiceServer
	rs.UnimplementedReplicationServiceServer
	peers             map[string]rs.ReplicationServiceClient // map of peer addresses to clients
	peerLock          sync.RWMutex
	leader            rs.ReplicationServiceClient
	Highest_Bid       int64
	Highest_BidId     int64
	Our_Timestamp     int64
	Highest_Timestamp int64
	IsLeader          bool
	peerPorts         []string
}


//Takes a bid from a bidder
//If not leader, it propagates the bid to the leader
//If leader, it updates the highest bid and propagates the bid to followers
func (n *P2PNode) Bid(ctx context.Context, bid *as.Amount) (ack *as.Ack, err error) {
	if !n.IsLeader {
		response, _ := n.leader.ReplicateBid(ctx, &rs.NewBid{
			Amount: bid.Amount,
			Bidderid:     bid.Bidderid,
		})

		if (response.Ack) {
			return &as.Ack{
				Ack: response.Ack,
				Bidderid:  response.Bidderid,
			}, nil
		}
	}

	response := n.ReplicateBid(ctx, &rs.NewBid{
		Amount: bid.Amount,
		Bidderid:     bid.Bidderid,
	})

	return &as.Ack{
		Ack: response.Ack,
		Bidderid:  response.Bidderid,
	}, nil
}

func (n *P2PNode) ConfirmLeader(NewLeader rs.NewLeader) {

}

func (n *P2PNode) result() (outcome int64) {
	//return highest value
	return n.Highest_Bid
}

func (n *P2PNode) ReplicateBid(ctx context.Context, bid *rs.NewBid) (ack *rs.Response) {
	if bid.Amount > n.Highest_Bid {
		n.UpdateFollowers(bid)
		n.Highest_Bid = bid.Amount
		n.Highest_BidId = bid.Bidderid
		n.Highest_Timestamp = n.Highest_Timestamp + 1
		return &rs.Response{
			Ack: true,
			Bidderid:  bid.Bidderid, //ADD CHECK TO SEE IF ID IS 0
		}
	}
	return &rs.Response{
		Ack: false,
		Bidderid:  bid.Bidderid,
	}
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

func startServer(node *P2PNode, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuctionServer(grpcServer, node)

	log.Printf("P2P node is running on port %s/n", address)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	if len(node.peers) == 0 {
		//set node leader to be the leader
		node.IsLeader = true
	}

	//put the node into the peers map

}
