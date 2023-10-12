package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "chatapp/proto"
)

const (
	address = "localhost:50051"
	//defaultUser = "Alice"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatappServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := c.ReceiveMessage(ctx)
	if err != nil {
		log.Fatalf("could not receive message: %v", err)
	}
	//Register User
	fmt.Print("Enter your user name:")
	var user string
	fmt.Scanln(fmt.Scanln(&user))
	// Send an initial "join" message
	err = stream.Send(&pb.Message{User: user, Text: "join"})
	if err != nil {
		log.Fatalf("failed to send join message: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		defer close(waitc)
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Printf("failed to receive message: %v", err)
				return
			}
			fmt.Printf("%s: %s\n", in.User, in.Text)
		}
	}()

	for {
		fmt.Print("Enter your message: ")
		var text string
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "quit" {
			break
		}
		msg := &pb.Message{User: user, Text: text}
		err = stream.Send(msg)
		if err != nil {
			log.Printf("failed to send message: %v", err)
			continue
		}
	}

	if err := stream.CloseSend(); err != nil {
		log.Printf("failed to close stream: %v", err)
	}

	select {
	case <-waitc:
	case <-ctx.Done():
	}
}
