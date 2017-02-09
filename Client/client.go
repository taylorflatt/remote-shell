package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/taylorflatt/Lab1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRemoteCommandClient(conn)

	reader := bufio.NewReader(os.Stdin)
	tCmd, _ := reader.ReadString('\n')
	tCmd2 := strings.Split(tCmd, " ")

	fmt.Printf("Name: %s", tCmd2[0])
	fmt.Printf("Args: %v", tCmd2[1:])

	cmdName := tCmd2[0]
	cmdArgs := []string{}

	if len(tCmd2) > 1 {
		temp := strings.Split(tCmd2[len(tCmd2)-1], "\n")
		cmdArgs[len(tCmd2)-1] = temp[0]
	} else {
		temp := strings.Split(tCmd2[len(tCmd2)-1], "\n")
		cmdName = temp[0]
	}

	r, err := c.SendCommand(context.Background(), &pb.CommandRequest{CmdName: cmdName, CmdArgs: cmdArgs})
	if err != nil {
		log.Fatalf("Command failed: %v", err)
	}

	log.Printf("%s", r.Output)
}
