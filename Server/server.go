package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"

	pb "github.com/taylorflatt/Lab1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// Executes a remote command from a client and returns that output. Otherwise, it will print an error.
func executeCommand(commandName string, commandArgs []string) string {
	tOutput, err := exec.Command(commandName, commandArgs...).Output()
	output := string(tOutput)

	fmt.Printf("Client ran: %s\n", commandName)
	fmt.Printf("Client with args: %v\n", commandArgs)
	fmt.Printf("Output for the command was: %s\n", output)

	if err == nil {
		return output
	}

	fmt.Println(err)
	return "An error was discovered executing your command."
}

// SayHello implements helloworld.GreeterServer
func (s *server) SendCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandReply, error) {

	var cmdName = in.CmdName
	var cmdArgs = in.CmdArgs
	var output = executeCommand(cmdName, cmdArgs)

	return &pb.CommandReply{Output: output}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRemoteCommandServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
