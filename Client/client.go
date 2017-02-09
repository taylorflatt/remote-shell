package main

import (
	"bufio"
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

	// Close the connection after main returns.
	defer conn.Close()

	// Create the client
	c := pb.NewRemoteCommandClient(conn)

	// Read in the user's command.
	r := bufio.NewReader(os.Stdin)
	tCmd, _ := r.ReadString('\n')
	tCmd2 := strings.Split(tCmd, " ")

	// Parse their input.
	cmdName := tCmd2[0]
	cmdArgs := []string{}

	// Strip off trailing carriage returns
	if len(tCmd2) > 1 {
		temp := strings.Split(tCmd2[len(tCmd2)-1], "\n")
		cmdArgs[len(tCmd2)-1] = temp[0]
	} else {
		temp := strings.Split(tCmd2[len(tCmd2)-1], "\n")
		cmdName = temp[0]
	}

	// Gets the response of the shell command from the server.
	res, err := c.SendCommand(context.Background(), &pb.CommandRequest{CmdName: cmdName, CmdArgs: cmdArgs})

	if err != nil {
		log.Fatalf("Command failed: %v", err)
	}

	log.Printf("%s", res.Output)
}
