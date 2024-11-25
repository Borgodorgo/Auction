package main

import (
	as "Replication/AuctionService"
	rs "Replication/ReplicationService"
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type P2PNode struct {
	as.UnimplementedAuctionServiceServer
	rs.UnimplementedReplicationServiceServer
	peers             map[string]rs.ReplicationServiceClient // map of peer addresses to clients
	peerLock          sync.RWMutex
	leader            rs.ReplicationServiceClient
	Highest_Bid       int64
	Highest_BidderId  int64
	Our_Timestamp     int64
	Highest_Timestamp int64
	IsLeader          bool
	peerPorts         []string
	active            bool
	counter           int
	address           string
	countdown         int
}

// Takes a bid from a bidder
// If not leader, it propagates the bid to the leader
// If leader, it updates the highest bid and propagates the bid to followers
func (n *P2PNode) Bid(ctx context.Context, bid *as.Amount) (ack *as.Ack, err error) {
	if !n.IsLeader && n.active {
		log.Printf("Server %s propagating to leader", n.address)
		response, err := n.leader.PropagateToLeader(ctx, &rs.NewBid{
			Amount:   bid.Amount,
			Bidderid: bid.Bidderid,
		})
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return &as.Ack{
				Ack:      false,
				Bidderid: bid.Bidderid,
			}, nil
		}
		return &as.Ack{
			Ack:      response.Ack,
			Bidderid: response.Bidderid,
		}, nil
	}

	if bid.Bidderid == 0 {
		n.Highest_BidderId++
		bid.Bidderid = n.Highest_BidderId
	}
	newBid := &rs.NewBid{
		Amount:   bid.Amount,
		Bidderid: bid.Bidderid,
	}

	if n.CheckBidValidity(newBid) {
		log.Printf("Bidder %d Successfully bidded %d", newBid.Bidderid, newBid.Amount)
		newBid = n.UpdateBidAsLeader(newBid)
		n.UpdateFollowers(newBid)
		return &as.Ack{
			Ack:      true,
			Bidderid: bid.Bidderid,
		}, nil
	}

	return &as.Ack{
		Ack:      false,
		Bidderid: bid.Bidderid,
	}, nil
}

func (n *P2PNode) CheckBidValidity(bid *rs.NewBid) (valid bool) {
	log.Printf("Checking bid validity Bid: %d Highest: %d", bid.Amount, n.Highest_Bid)
	return bid.Amount > n.Highest_Bid
}

func (n *P2PNode) Result(ctx context.Context, empty *emptypb.Empty) (result *as.Outcome, err error) {
	return &as.Outcome{
		Amount: n.Highest_Bid,
	}, err
}

func (n *P2PNode) UpdateBidAsLeader(bid *rs.NewBid) (newbid *rs.NewBid) {
	n.Highest_Bid = bid.Amount
	n.Highest_Timestamp = n.Highest_Timestamp + 1
	bidderid := bid.Bidderid
	if bid.Bidderid == 0 {
		n.Highest_BidderId++
		bidderid = n.Highest_BidderId
	}
	return &rs.NewBid{
		Amount:    bid.Amount,
		Bidderid:  bidderid,
		Timestamp: n.Highest_Timestamp,
	}
}

// Run only by leader node
func (n *P2PNode) PropagateToLeader(ctx context.Context, bid *rs.NewBid) (ack *rs.Response, err error) {
	log.Printf("Leader received bid from %d", bid.Bidderid)
	if n.CheckBidValidity(bid) {
		bid = n.UpdateBidAsLeader(bid)
		n.UpdateFollowers(bid)
		return &rs.Response{
			Ack:      true,
			Bidderid: n.Highest_BidderId,
		}, nil
	}
	return &rs.Response{
		Ack:      false,
		Bidderid: bid.Bidderid,
	}, nil
}

// Run only by follower nodes
func (n *P2PNode) ReplicateBid(ctx context.Context, bid *rs.NewBid) (ack *rs.Response, err error) {
	log.Printf("Follower %s received bid from %d", n.address, bid.Bidderid)
	n.Highest_Bid = bid.Amount
	n.Highest_BidderId = bid.Bidderid
	n.Highest_Timestamp = bid.Timestamp

	return &rs.Response{
		Ack: true,
	}, err
}

func (n *P2PNode) UpdateFollowers(newbid *rs.NewBid) {
	for address := range n.peers {
		if peer, exists := n.peers[address]; exists {
			_, err := peer.ReplicateBid(context.Background(), newbid)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	}
}

func (n *P2PNode) ConfirmLeader(ctx context.Context, NewLeader *rs.NewLeader) (response *rs.Response, err error) {

	log.Printf("New leader is %s", NewLeader.Address)
	address := "localhost" + NewLeader.Address
	n.leader = n.peers[address]
	n.IsLeader = false

	return &rs.Response{
		Ack: true,
	}, err
}

func (n *P2PNode) HeartBeat(ctx context.Context, empty *emptypb.Empty) (response *rs.Response, err error) {
	//if heartbeat received, reset timeout
	n.counter = 0
	return &rs.Response{
		Ack: true,
	}, err
}

func (n *P2PNode) HeartBeating() {
	//send heartbeat to all peers
	for {
		if n.active {
			log.Println("Heartbeat")
			for address := range n.peers {
				if address != "localhost"+n.address {
					if peer, exists := n.peers[address]; exists {
						_, err := peer.HeartBeat(context.Background(), &emptypb.Empty{})
						if err != nil {
							log.Printf("Error sending message: %v", err)
						}
					}
				}
			}
			time.Sleep(1 * time.Second)
			n.countdown++
			if n.countdown > 100 {
				//n.active = false
				//stopAuction()
			}
		}
	}
}

func (n *P2PNode) Election() {
	//send election message to all peers
	log.Println("Election time")
	n.peerLock.RLock()
	counter := len(n.peers) - 1
	for address := range n.peers {
		if peer, exists := n.peers[address]; exists {
			ack, err := peer.ConfirmLeader(context.Background(), &rs.NewLeader{
				Address: n.address,
			})
			if err != nil {
				log.Print("Heartbeat failed")
				log.Printf("Error sending message: %v", err)
				delete(n.peers, address)
				break
			}
			if ack.Ack {
				counter--
			}
			if counter == 0 {
				break
			}
		}

	}
	n.IsLeader = true
	n.active = true
	n.counter = 0
	go n.HeartBeating()
	n.peerLock.RUnlock()
}

func createServer(node *P2PNode, address string, leader bool) {

	node.IsLeader = leader
	node.address = address
	lis, err := net.Listen("tcp", "localhost"+address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	as.RegisterAuctionServiceServer(grpcServer, node)
	rs.RegisterReplicationServiceServer(grpcServer, node)

	log.Printf("P2P node is running on port %s/n", address)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
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
		active:            true,
	}
	return node
}

func (n *P2PNode) Crash() {
	log.Println("Is crashing")
	n.active = false
	n.peerLock.Lock()
}

func (n *P2PNode) PeerSetup() {
	n.peerPorts[0] = ":5001"
	n.peerPorts[1] = ":5002"
	// n.peerPorts[2] = ":5003"
	// n.peerPorts[3] = ":5004"
	// n.peerPorts[4] = ":5005"

	for i := 0; i < len(n.peerPorts); i++ {
		if n.address != n.peerPorts[i] {
			address := "localhost" + n.peerPorts[i]
			conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("Failed to connect to peer %s: %v", address, err)
				return
			}

			client := rs.NewReplicationServiceClient(conn)

			if i == 0 {
				n.leader = client
			}
			n.peerLock.Lock()
			n.peers[address] = client
			n.peerLock.Unlock()
		}
	}
}

func (n *P2PNode) startServer() {
	for {
		log.Printf("Counter: %d", n.counter)
		n.counter++
		if n.counter > 5 && n.address == ":5002" {
			n.Election()
			break

		}
		time.Sleep(1 * time.Second)
	}
}

func (n *P2PNode) startLeader() {
	log.Print("Leader starting heartbeat")
	go n.HeartBeating()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	leader := false
	if input == ":5001" {
		log.Println("Leader")
		leader = true
	}
	node := CreateNode()
	node.PeerSetup()
	go createServer(node, input, leader)

	if leader {
		node.startLeader()
	} else {
		node.startServer()
	}

	for {
		time.Sleep(time.Second)
	}
}
