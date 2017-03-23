package server

import (
	"log"
	"net"
	"os/exec"

	pb "github.com/taylorflatt/remote-shell"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// The port the server is listening on.
const (
	port = ":12021"
)

// Server is used to implement the RemoteCommandServer
type server struct{}

// ExecuteCommand actually runs the command.
// It returns the output of whichever command was run.
func ExecuteCommand(commandName string, commandArgs []string) string {
	tOutput, err := exec.Command(commandName, commandArgs...).Output()
	output := string(tOutput)

	if err != nil {
		return err.Error()
	}

	return output
}

// SendCommand receives the command from the client and then executes it server-side.
// It returns a commmand reply consisting of the output of the command.
func (s *server) SendCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandReply, error) {

	var cmdName = in.CmdName
	var cmdArgs = in.CmdArgs
	var output = ExecuteCommand(cmdName, cmdArgs)

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
