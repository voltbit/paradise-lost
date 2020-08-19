package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "./pingpb"

	"google.golang.org/grpc"
)

var id string
var wg sync.WaitGroup

func generateId(length int) string {
	if length < 1 {
		length = 10
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	id := make([]rune, length)
	var letter rune
	for i := 0; i < length; i++ {
		if rand.Int()%2 == 0 {
			letter = 'a'
		} else {
			letter = 'A'
		}
		id[i] = letter + rune(rand.Intn(26))
	}
	return string(id)
}

func sendPing() {
	conn, err := grpc.Dial("localhost:50059", grpc.WithInsecure(), grpc.WithBlock())
	fmt.Println("Connection made")
	if err != nil {
		fmt.Println("Could not open connection")
	}
	defer conn.Close()

	c := pb.NewPingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Sending ping")
	r, err := c.SendPing(ctx, &pb.PingRequest{Id: id})
	if err != nil {
		fmt.Println("Could not send ping")
	}
	fmt.Println("Received pong from", r.Id)
}

type server struct {
	pb.UnimplementedPingServer
}

func startPingListener() {
	go func() {
		fmt.Println("Starting sync gRPC server")
		lis, err := net.Listen("tcp", "localhost:50059")
		if err != nil {
			fmt.Println("Could not start listener", err)
		}
		s := grpc.NewServer()
		pb.RegisterPingServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			fmt.Println("Error while serving", err)
		}
		fmt.Println("Sync gRPC server exiting")
		wg.Done()
	}()
}

func (s *server) SendPing(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Id: id}, nil
}

func schedule(f func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				f()
			}
		}
	}()
	return ticker
}

func main() {
	wg.Add(2)
	id = generateId(5)
	startPingListener()
	schedule(sendPing, 2*time.Second)
	wg.Wait()
}
