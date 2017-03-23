package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/taylorflatt/remote-shell"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":12021"
)

func main() {
	// Read in the user's command.
	r := bufio.NewReader(os.Stdin)

	// Read the server address
	fmt.Print("Please specify the server IP: ")
	address, _ := r.ReadString('\n')
	address = strings.TrimSpace(address)
	address = address + port

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Close the connection after main returns.
	defer conn.Close()

	// Create the client
	c := pb.NewRemoteCommandClient(conn)

	fmt.Printf("\nYou have successfully connected to %s! To disconnect, hit ctrl+c or type exit.\n", address)

	// Keep connection alive until ctrl+c or exit is entered.
	for true {
		fmt.Print("$ ")
		tCmd, _ := r.ReadString('\n')

		// This strips off any trailing whitespace/carriage returns.
		tCmd = strings.TrimSpace(tCmd)
		tCmd2 := strings.Split(tCmd, " ")

		// Parse their input.
		cmdName := tCmd2[0]

		//cmdArgs := []string{}
		cmdArgs := tCmd2[1:]

		// Close the connection if the user enters exit.
		if cmdName == "exit" {
			break
		}

		// Gets the response of the shell comm and from the server.
		res, err := c.SendCommand(context.Background(), &pb.CommandRequest{CmdName: cmdName, CmdArgs: cmdArgs})

		if err != nil {
			log.Fatalf("Command failed: %v", err)
		}

		log.Printf("    %s", res.Output)
	}
}
