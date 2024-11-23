package main

import "time"

/*
	func oldmain() {
		node1 := CreateNode(0, 3)
		node2 := CreateNode(1, 3)
		node3 := CreateNode(2, 3)

		// Start the gRPC server
		go StartServer(node1, ":50051")
		go StartServer(node2, ":50052")
		go StartServer(node3, ":50053")

		time.Sleep(3 * time.Second)

		// Add peers (simulate peer discovery for demonstration)
		node1.AddPeer("localhost:50052", 1)
		node1.AddPeer("localhost:50053", 2)

		node2.AddPeer("localhost:50051", 0)
		node2.AddPeer("localhost:50053", 2)

		node3.AddPeer("localhost:50051", 0)
		node3.AddPeer("localhost:50052", 1)

		go node1.Start()
		go node2.Start()
		go node3.Start()

		counter := 0
		for {
			if counter > 100 {
				break
			}
			time.Sleep(time.Second)
			counter++
		}
	}
*/
func main() {
	nodes := make([]*P2PNode, 5)
	for i := 0; i < 5; i++ {
		nodes[i] = CreateNode()
	}

	for i := 0; i < 5; i++ {
		nodes[i].PeerSetup()
	}

	leader := true
	for i := 0; i < 5; i++ {

		go createServer(nodes[i], nodes[i].peerPorts[i], leader)
		leader = false
	}

	for i := 0; i < 5; i++ {
		go nodes[i].startServer()
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < 5; i++ {
		go start()
	}

	time.Sleep(3 * time.Second)
	go nodes[0].Crash()
	for {
		time.Sleep(time.Second)
	}

}
