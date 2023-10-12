package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "chatapp/proto"
)

type server struct {
	pb.UnimplementedChatappServiceServer
	mu      sync.Mutex
	clients map[string]pb.ChatappService_ReceiveMessageServer
}

func (s *server) SendMessage(ctx context.Context, in *pb.Message) (*empty.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("%s: %s\n", in.User, in.Text)
	var lastError error
	for user, client := range s.clients {
		if user != in.User {
			err := client.Send(in)
			if err != nil {
				lastError = err
				delete(s.clients, user) // Remove the disconnected client
				continue                // Continue with the next client
			}
		}
	}
	return &emptypb.Empty{}, lastError
}

func (s *server) ReceiveMessage(stream pb.ChatappService_ReceiveMessageServer) error {
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.clients[msg.User] = stream
	s.mu.Unlock()
	fmt.Printf("%s joined the chat\n", msg.User)
	for {
		msg, err := stream.Recv()
		if err != nil {
			s.mu.Lock()
			if msg != nil { // Check if msg is not nil before accessing its members
				delete(s.clients, msg.User)
				fmt.Printf("%s left the chat\n", msg.User)
			}
			s.mu.Unlock()
			return err
		}
		fmt.Printf("%s: %s\n", msg.User, msg.Text)
		_, err = s.SendMessage(stream.Context(), msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatappServiceServer(s, &server{clients: make(map[string]pb.ChatappService_ReceiveMessageServer)})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
