package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/calculi-corp/config"
	client "github.com/calculi-corp/grpc-client"
	"github.com/calculi-corp/log"
	pb "github.com/scline12/grEcho/pb"
)

func main() {
	// Read config and set logging level
	err := config.Config.SetCliFlags()
	if log.CheckErrorf(err, "Error reading configuration") {
		os.Exit(1)
	}
	fmt.Println("Flag Values:\n", config.Config.FlagValues())
	logLvl := config.Config.GetString("logging.level")
	if logLvl != "" {
		log.SetLoggingLevel(logLvl)
	}

	// Instantiate client
	clt, err := client.NewClient()
	if err != nil {
		log.Error("Failed to set up Client", err)
		return
	}
	defer clt.Close() // Clean up client

	// Create host address. For most GR services, this would be pulled from the config.
	host := fmt.Sprintf("127.0.0.1:%d", config.Config.GetInt("server.listen.port"))

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter text to echo:")
	fmt.Print("\nRequest: ") // prompt
	for scanner.Scan() {
		req := &pb.EchoRequest{
			Message: scanner.Text(),
		}

		rsp := &pb.EchoResponse{}

		err = clt.SendGrpc(host, pb.GetDesc().ServiceName, "Echo", req, rsp) // Send the request
		if err != nil {
			log.Error("Failed to get a response from server", err)
		}

		fmt.Printf("Response: %s\n", rsp.GetMessage()) // print reply
		fmt.Print("\nRequest: ")                       // prompt
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
