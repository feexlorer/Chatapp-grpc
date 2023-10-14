package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "chatapp/proto"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatappServiceClient(conn)

	streamCtx, streamCancel := context.WithCancel(context.Background())
	defer streamCancel()

	stream, err := c.ReceiveMessage(streamCtx)
	if err != nil {
		log.Fatalf("could not receive message: %v", err)
	}

	a := app.New()
	loginWindow := a.NewWindow("Login")
	loginWindow.Resize(fyne.Size{Width: 300, Height: 200})

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter your username")

	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		err = stream.Send(&pb.Message{User: username, Text: "join"})
		if err != nil {
			log.Fatalf("failed to send join message: %v", err)
		}

		loginWindow.Close()

		chatWindow := a.NewWindow("Chat App")
		chatWindow.Resize(fyne.Size{Width: 600, Height: 400})

		input := widget.NewEntry()
		input.SetPlaceHolder("Enter message")

		display := widget.NewLabel("")

		sendButton := widget.NewButton("Send", func() {
			msg := &pb.Message{User: username, Text: input.Text}
			err = stream.Send(msg)
			if err != nil {
				log.Printf("failed to send message: %v", err)
				return
			}
			input.SetText("")
		})

		usersButton := widget.NewButton("Users", func() {
			msg := &pb.Message{User: username, Text: "userinfo"}
			err = stream.Send(msg)
			if err != nil {
				log.Printf("failed to send userinfo message: %v", err)
				return
			}
		})

		content := container.NewVBox(display, input, container.NewHBox(sendButton, usersButton))

		go func() {
			for {
				in, err := stream.Recv()
				if err != nil {
					log.Printf("failed to receive message: %v", err)
					return
				}
				display.SetText(display.Text + "\n" + in.User + ": " + in.Text)
				display.Refresh()
			}
		}()

		chatWindow.SetContent(content)
		chatWindow.Show()
		chatWindow.SetOnClosed(func() {
			streamCancel()
		})
	})

	loginContent := container.NewVBox(usernameEntry, loginButton)
	loginWindow.SetContent(loginContent)

	loginWindow.ShowAndRun()
}
