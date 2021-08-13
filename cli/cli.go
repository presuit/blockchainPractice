package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/presuit/nomadcoin/explorer"
	"github.com/presuit/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags\n\n")
	fmt.Printf("-port=4000	:	Set the PORT of the server\n")
	fmt.Printf("-mode=rest	:	Start html or rest\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "set the port")
	mode := flag.String("mode", "rest", "choose between 'html' and 'rest'")

	flag.Parse()

	fmt.Println(*port, *mode)

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	case "both":
		go explorer.Start(*port + 1)
		rest.Start(*port)
	default:
		usage()

	}
}
