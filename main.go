package main

import (
	"fmt"
	"os"

	"github.com/calculi-corp/config"
	client "github.com/calculi-corp/grpc-client"
	server "github.com/calculi-corp/grpc-server"
	"github.com/calculi-corp/log"
	echo "github.com/scline12/grEcho/handler"
)

func main() {
	err := config.Config.SetCliFlags()
	if log.CheckErrorf(err, "Error reading configuration") {
		os.Exit(1)
	}
	fmt.Println("Flag Values:\n", config.Config.FlagValues())
	logLvl := config.Config.GetString("logging.level")
	if logLvl != "" {
		log.SetLoggingLevel(logLvl)
	}
	server, err := server.NewServer()
	if log.CheckErrorf(err, "Unable to instantiate server") {
		os.Exit(1)
	}
	defer server.Stop() // Clean up when you are done

	clt, err := client.NewClient() // Creating a new client
	if log.CheckErrorf(err, "Failed to create client") {
		os.Exit(2)
	}

	handler := echo.NewEchoHandler(clt)
	// Pass the certificate and key files to the server
	err = server.AddHandler(handler) //Register a service
	if log.CheckErrorf(err, "Failure adding handler") {
		os.Exit(3)
	}

	server.Start()
	server.WaitForExit() // Blocks until the server exits cleanly on request
}
