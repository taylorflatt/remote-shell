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
	r := bufio.NewReader(os.Stdin)

	/* Read the server address
	address, _ := r.ReadString('\n')

	temp := strings.Split(address, "\n")
	address = strings.Join(temp[:], "")
	*/

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Close the connection after main returns.
	defer conn.Close()

	// Create the client
	c := pb.NewRemoteCommandClient(conn)

	// Read in the user's command.
	//r := bufio.NewReader(os.Stdin)

	for true {
		tCmd, _ := r.ReadString('\n')
		tCmd2 := strings.Split(tCmd, " ")

		// Parse their input.
		cmdName := tCmd2[0]
		cmdArgs := []string{}

		// Strip off trailing carriage returns
		if len(tCmd2) > 1 {
			cmdArgs = tCmd2[1:]
			fmt.Printf("Args: %v", cmdArgs)
			//temp := strings.TrimRight(tCmd2[len(tCmd2)-1], "\n")

			fmt.Printf("Temp: %v", cmdArgs)
			fmt.Printf("Length: %d", len(tCmd2)-1)

			cmdArgs[len(cmdArgs)-1] = strings.TrimRight(cmdArgs[len(cmdArgs)-1], "\n")
		} else {
			temp := strings.TrimRight(tCmd2[len(tCmd2)-1], "\n")
			cmdName = temp
		}

		// Close the connection if the user enters exit.
		if cmdName == "exit" {
			break
		}

		// Gets the response of the shell command from the server.
		res, err := c.SendCommand(context.Background(), &pb.CommandRequest{CmdName: cmdName, CmdArgs: cmdArgs})

		if err != nil {
			log.Fatalf("Command failed: %v", err)
		}

		log.Printf("%s", res.Output)
	}
}
