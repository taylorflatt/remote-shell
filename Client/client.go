package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	pb "github.com/taylorflatt/Lab1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "192.168.0.170:50051"
)

func main() {
	// Read in the user's command.
	r := bufio.NewReader(os.Stdin)

	/* Read the server address
	address, _ := r.ReadString('\n')

	temp := strings.Split(address, "\n")
	address = strings.Join(temp[:], "")
	*/

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Close the connection after main returns.
	defer conn.Close()

	// Create the client
	c := pb.NewRemoteCommandClient(conn)

	// Keep connection alive until ctrl+c or exit is entered.
	for true {
		tCmd, _ := r.ReadString('\n')
		tCmd2 := strings.Split(tCmd, " ")

		// Parse their input.
		cmdName := tCmd2[0]
		cmdArgs := []string{}

		// Strip off trailing carriage returns
		if len(tCmd2) > 1 {
			cmdArgs = tCmd2[1:]
			if runtime.GOOS == "windows" {
				cmdArgs[len(cmdArgs)-1] = strings.TrimRight(cmdArgs[len(cmdArgs)-1], "\r")
			} else {
				cmdArgs[len(cmdArgs)-1] = strings.TrimRight(cmdArgs[len(cmdArgs)-1], "\n")
			}
		} else {
			if runtime.GOOS == "windows" {
				fmt.Printf("AFTER %v", cmdName)
				cmdName = strings.TrimRight(cmdName, "\r")
				cmdName = strings.TrimRight(cmdName, "\n")
			} else {
				cmdName = strings.TrimRight(cmdName, "\n")
			}
		}

		// Close the connection if the user enters exit.
		if cmdName == "exit" {
			break
		}

		// Gets the response of the shell comm and from the server.
		res, err := c.SendCommand(context.Background(), &pb.CommandRequest{CmdName: cmdName, CmdArgs: cmdArgs})

		if err != nil {
			log.Fatalf("Command failed: %v", err)
		}

		log.Printf("%s", res.Output)
	}
}
