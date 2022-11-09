package main

import (
	"fmt"

	"github.com/kumarabd/CS6378-Project2/go/config"
	"github.com/kumarabd/CS6378-Project2/go/internal/node"
	"github.com/kumarabd/CS6378-Project2/go/logger"
)

var (
	ID string = "0"
)

func main() {
	// Initialize Logger instance
	log, err := logger.New("Node-"+ID, logger.Options{
		Format:     logger.SyslogLogFormat,
		DebugLevel: true,
	})
	if err != nil {
		fmt.Println("Logger failed")
	}
	log.Info("logger initialized")

	cfg, err := config.ReadConfig("config.txt")
	if err != nil {
		log.Error(err)
	}
	log.Info("config initialized")

	// Create node
	nodeObj, err := node.New(ID, cfg, log)
	if err != nil {
		log.Error(err)
	}
	log.Info("created node")

	// Start node
	log.Info("running")
	nodeObj.Start()
	log.Info("exiting")
}
