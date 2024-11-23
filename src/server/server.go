package main

import (
	pb "assignment5/auction/src/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
)

const (
	RoleLeader   = "leader"
	RoleFollower = "follower"
)

type Server struct {
	pb.UnimplementedAuctionServiceServer

	ID         string   // Unique ID of the server
	Role       string   // "leader" or "follower"
	LeaderID   string   // ID of the current leader
	Peers      []string // List of other server addresses
	Bids       []pb.BidRequest
	HighestBid *pb.BidRequest

	Mutex sync.Mutex // To protect shared state
}

func (s *Server) Bid(ctx context.Context, req *pb.BidRequest) (*pb.BidResponse, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.Role == RoleFollower {
		// Forward the bid to the leader
		return forwardBidToLeader(s.LeaderID, req)
	}

	// Check if bid is higher
	if s.HighestBid == nil || req.Amount > s.HighestBid.Amount {
		s.HighestBid = req
		s.Bids = append(s.Bids, *req)

		// Sync with followers
		for _, peer := range s.Peers {
			go func(peer string) {
				if err := syncBidWithFollower(peer, req); err != nil {
					log.Printf("Failed to sync with follower %s: %v", peer, err)
				}
			}(peer)
		}

		return &pb.BidResponse{
			Success: true,
			Message: "Bid accepted",
		}, nil
	}

	return &pb.BidResponse{
		Success: false,
		Message: "Bid too low",
	}, nil
}

func (s *Server) GetStatus(ctx context.Context, _ *pb.StatusRequest) (*pb.StatusResponse, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.HighestBid == nil {
		return &pb.StatusResponse{
			HighestBidder: "",
			HighestBid:    0,
			Message:       "No bids yet",
		}, nil
	}

	return &pb.StatusResponse{
		HighestBidder: s.HighestBid.Bidder,
		HighestBid:    s.HighestBid.Amount,
		Message:       "Current highest bid",
	}, nil
}

func (s *Server) SyncBid(ctx context.Context, req *pb.BidRequest) (*pb.SyncResponse, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// Update the follower's state
	if s.HighestBid == nil || req.Amount > s.HighestBid.Amount {
		s.HighestBid = req
		s.Bids = append(s.Bids, *req)
	}

	return &pb.SyncResponse{
		Success: true,
		Message: "Bid synced successfully",
	}, nil
}

func syncBidWithFollower(followerAddr string, bid *pb.BidRequest) error {
	conn, err := grpc.NewClient(followerAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to follower: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuctionServiceClient(conn)
	_, err = client.SyncBid(context.Background(), bid)
	return err
}

func forwardBidToLeader(leaderAddr string, bid *pb.BidRequest) (*pb.BidResponse, error) {
	conn, err := grpc.NewClient(leaderAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to leader: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuctionServiceClient(conn)
	return client.Bid(context.Background(), bid)
}

func RunServer(s *Server, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuctionServiceServer(grpcServer, s)

	log.Printf("Server %s running as %s on port %s", s.ID, s.Role, port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: go run main.go <serverID> <role> <port>")
	}

	serverID := os.Args[1]
	role := os.Args[2]
	port := os.Args[3]

	if role != RoleLeader && role != RoleFollower {
		log.Fatalf("Invalid role: must be 'leader' or 'follower'")
	}

	// Example peer addresses, excluding the current server's address
	peers := []string{":50051", ":50052", ":50053"}
	peers = remove(peers, port)

	s := &Server{
		ID:       serverID,
		Role:     role,
		LeaderID: "", // Set if role is follower
		Peers:    peers,
		Bids:     []pb.BidRequest{},
	}

	if role == RoleLeader {
		s.LeaderID = port
	} else {
		// Assume the first peer is the leader
		s.LeaderID = peers[0]
	}

	RunServer(s, port)
}

func remove(peers []string, addr string) []string {
	var result []string
	for _, peer := range peers {
		if peer != addr {
			result = append(result, peer)
		}
	}
	return result
}
