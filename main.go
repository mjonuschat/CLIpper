package main

import (
	"clipper/jsonrpcclient"
	"clipper/ui"
	"fmt"
	"log"
	"os"
	"strings"
)

var version = "?"
var commit = ""

func StartInteractive(url string) {
	versionString := version
	if commit != "" {
		versionString = versionString + "-" + commit[:7]
	}
	rpcClient := jsonrpcclient.NewClient(url)
	defer rpcClient.Close()

	if err := rpcClient.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	tui := ui.NewTUI(rpcClient, versionString)
	if _, err := tui.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]
	var url string
	switch len(args) {
	case 2:
		url = "ws://" + args[0] + ":" + string(args[1]) + "/websocket"
	case 1:
		if strings.Contains(args[0], "://") {
			url = args[0]
		} else {
			url = "ws://" + args[0] + "/websocket"
		}
	default:
		fmt.Println("clipper " + version + "-" + commit)
		fmt.Println("\nUsage: " +
			"\n\t" + os.Args[0] + " <hostname/ip> [<port>]" +
			"\n\t" + os.Args[0] + " <hostname/ip>:<port>" +
			"\n\t" + os.Args[0] + " ws[s]://<hostname/ip>[:<port>]/path")
		fmt.Println("\nExamples: " +
			"\n\t" + os.Args[0] + " mainsailos.local" +
			"\n\t" + os.Args[0] + " ws://mainsailos.local/websocket" +
			"\n\t" + os.Args[0] + " myvzero.local 6969" +
			"\n\t" + os.Args[0] + " 192.168.1.69:8000\n")
		os.Exit(1)

		url = "ws://localhost/websocket"
	}

	StartInteractive(url)
}
