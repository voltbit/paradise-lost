package cluster

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "../memdbpb"
	"google.golang.org/grpc"
)

type MemdbClusterComms struct {
	id     string
	addr   string
	server *grpc.Server
}

var nodeId string

func NewMemdbClusterComms(id string, url string) *MemdbClusterComms {
	nodeId = id
	return &MemdbClusterComms{
		id:     id,
		addr:   url,
		server: grpc.NewServer(),
	}
}

func (t *MemdbClusterComms) sendPing(nodeList [][]string) {
	fmt.Println("Sending to nodelist:", nodeList)
	for _, node := range nodeList {
		conn, err := grpc.Dial(node[1], grpc.WithInsecure(), grpc.WithBlock())
		// fmt.Println("Connection made")
		if err != nil {
			fmt.Println("Could not open connection")
		}
		defer conn.Close()

		c := pb.NewPingClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SendPing(ctx, &pb.PingRequest{Id: node[0]})
		if err != nil {
			fmt.Println("Could not send ping", err)
		}
		fmt.Println(t.id, "Received pong from", r.Id)
	}
}

type grpcProtobufServer struct {
	pb.UnimplementedPingServer
}

func (t *MemdbClusterComms) StartServer() {
	go func() {
		fmt.Println("Starting sync gRPC server")
		lis, err := net.Listen("tcp", t.addr)
		if err != nil {
			fmt.Println("Could not start listener", err)
		}
		s := grpc.NewServer()
		pb.RegisterPingServer(s, &grpcProtobufServer{})
		if err := s.Serve(lis); err != nil {
			fmt.Println("Error while serving", err)
		}
		fmt.Println("Sync gRPC server exiting on node", t.id)
	}()
}

func (s *grpcProtobufServer) SendPing(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Id: nodeId}, nil // TODO how to make this method use instance data?
}

func schedule(f func([][]string), nodeList [][]string, interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				f(nodeList)
			}
		}
	}()
	return ticker
}

func (t *MemdbClusterComms) StartPinging(interval time.Duration, nodeList [][]string) {
	schedule(t.sendPing, nodeList, interval)
}
