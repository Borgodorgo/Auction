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
	node1 := CreateNode()
	node2 := CreateNode()
	node3 := CreateNode()
	node4 := CreateNode()
	node5 := CreateNode()
	go startServer(node1, ":5001", true)
	go startServer(node2, ":5002", false)
	go startServer(node3, ":5003", false)
	go startServer(node4, ":5004", false)
	go startServer(node5, ":5005", false)

	time.Sleep(3 * time.Second)
	for i := 0; i < 5; i++ {
		go start()
	}

	time.Sleep(3 * time.Second)
	go node1.Crash()
	for {
		time.Sleep(time.Second)
	}

}
