package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jessevdk/go-flags"
	"github.com/lisuiheng/fsm/server"
)

// Build flags
var (
	version = "0.1.0"
)

func main() {
	srv := server.Server{
		Version: version,
	}

	parser := flags.NewParser(&srv, flags.Default)
	parser.ShortDescription = `FSM`
	parser.LongDescription = `Options for FSM`

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	if srv.ShowVersion {
		log.Printf("FSM %s\n", version)
		os.Exit(0)
	}

	ctx := context.Background()
	var g errgroup.WaitGroup
	if err := srv.Serve(ctx); err != nil {
		log.Fatalln(err)
	}
}
