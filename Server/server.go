package server

import (
	"log"
	"net"
	"os/exec"

	pb "github.com/taylorflatt/Lab1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":12021"
)

// Server is used to implement the RemoteCommandServer
type server struct{}

// Executes a remote command from a client and returns that output. Otherwise, it will print an error.
func executeCommand(commandName string, commandArgs []string) string {
	tOutput, err := exec.Command(commandName, commandArgs...).Output()
	output := string(tOutput)

	if err != nil {
		return err.Error()
	}

	return output
}

// This function executes the client's shell command and returns the results.
func (s *server) SendCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandReply, error) {

	var cmdName = in.CmdName
	var cmdArgs = in.CmdArgs
	var output = executeCommand(cmdName, cmdArgs)

	return &pb.CommandReply{Output: output}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// Initializes the gRPC server.
	s := grpc.NewServer()

	// Register the server with gRPC.
	pb.RegisterRemoteCommandServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
