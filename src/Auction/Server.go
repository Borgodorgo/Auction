package main

import (
	as "Replication/m/v2/AuctionService/Auction"
	rs "Replication/m/v2/ReplicationService/Replication"
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type P2PNode struct {
	as.UnimplementedAuctionServiceServer
	rs.UnimplementedReplicationServiceServer
	peers             map[string]rs.ReplicationServiceClient // map of peer addresses to clients
	peerLock          sync.RWMutex
	leader            rs.ReplicationServiceServer
	Highest_Bid       int64
	Highest_BidderId  int64
	Our_Timestamp     int64
	Highest_Timestamp int64
	IsLeader          bool
	peerPorts         []string
}

// Takes a bid from a bidder
// If not leader, it propagates the bid to the leader
// If leader, it updates the highest bid and propagates the bid to followers
func (n *P2PNode) Bid(ctx context.Context, bid *as.Amount) (ack *as.Ack, err error) {
	if !n.IsLeader {
		response, _ := n.leader.PropagateToLeader(ctx, &rs.NewBid{
			Amount: bid.Amount,
			Bidderid:     bid.Bidderid,
		})

		return &as.Ack{
			Ack: response.Ack,
			Bidderid:  response.Bidderid,
		}, nil
	}

	if CheckBidValidity(bid) {
		bid = n.UpdateBidAsLeader(bid)
		n.UpdateFollowers(bid)
		return &as.Ack{
			Ack: true,
			Bidderid: bid.Bidderid,
		}, nil
	}

	return &as.Ack{
		Ack: false,
		Bidderid:  bid.Bidderid,
	}, nil
}

func CheckBidValidity (bid *rs.NewBid) (valid bool) {
	if bid.Amount > n.Highest_Bid {
		return true
	}
	return false
}

func (n *P2PNode) Result(ctx context.Context, empty *emptypb.Empty) (result *as.Outcome, err error) {
	return &as.Outcome{
		Amount: n.Highest_Bid,
	}, err
}


func (n *P2PNode) UpdateBidAsLeader(bid *rs.NewBid) (bid *rs.NewBid) {
	n.Highest_Bid = bid.Amount
	n.Highest_Timestamp = n.Highest_Timestamp + 1
	bidderid := bid.Bidderid
	if bid.Bidderid == 0 {
		n.Highest_BidderId++
		bidderid = n.Highest_BidderId
	}
	return &rs.NewBid{
		Amount: bid.Amount,
		Bidderid:     bidderid,
		Timestamp: n.Highest_Timestamp,
	}
}

//Run only by leader node
func (n *P2PNode) PropagateToLeader(ctx context.Context, bid *rs.NewBid) (ack *rs.Response, err error) {
	if CheckBidValidity(bid) {
		bid = n.UpdateBidAsLeader(bid)
		n.UpdateFollowers(bid)
		return &rs.Response{
			Ack: true,
			Bidderid:  n.Highest_BidderId,
		}, nil
	}
	return &rs.Response{
		Ack: false,
		Bidderid:  bid.Bidderid,
	}, nil
}

//Run only by follower nodes
func (n *P2PNode) ReplicateBid(ctx context.Context, bid *rs.NewBid) (ack *rs.Response) {
		n.Highest_Bid = bid.Amount
		n.Highest_BidderId = bid.Bidderid
		n.Highest_Timestamp = bid.Timestamp

		return &rs.Response{
			Ack: true,
		}
}

func (n *P2PNode) UpdateFollowers(newbid *rs.NewBid) {
	n.peerLock.RLock()
	for address := range n.peers {
		if peer, exists := n.peers[address]; exists {
			_, err := peer.ReplicateBid(context.Background(), newbid)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	}
	n.peerLock.RUnlock()
}

func (n *P2PNode) ConfirmLeader(NewLeader *rs.NewLeader) (response *rs.Response) {
	n.leader = 
}

func (n *P2PNode) HeartBeat() (response *rs.Response) {
	//if heartbeat received, reset timeout
	
}

func (n *P2PNode) HeartBeatResponse() () {
	
}

func (n *P2PNode) result() (outcome int64) {
	//return highest value
	return n.Highest_Bid
}

func startServer(node *P2PNode, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	as.RegisterAuctionServiceServer(grpcServer, node)

	log.Printf("P2P node is running on port %s/n", address)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	if len(node.peers) == 0 {
		//set node leader to be the leader
		node.IsLeader = true
	}
	
	//if leader, go heartbeat function

	//put the node into the peers map

}

func CreateNode() *P2PNode {
	node := &P2PNode{
		peers:             make(map[string]rs.ReplicationServiceClient),
		Our_Timestamp:     1,
		Highest_Timestamp: 1,
		peerPorts:         make([]string, 5),
	}
	return node
}
