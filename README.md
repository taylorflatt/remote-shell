# Remote shell
A client-server remote shell implemented in Go using gRPC.

## Usage
* First, run the server by going to the Server directory and typing `go run server.go`.
* Next, run the client by going to the Client directory and typing `go run client.go`. 
* Finally, enter in the IP to the server and you will be able to type any commands you would like.

## Notes
To disconnect from the server, press ctrl+c or type exit (hit enter) and the client will disconnect from the server.

This client/server assumes a 12021 port.