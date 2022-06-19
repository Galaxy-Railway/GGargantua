package main

import (
	"flag"
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/cmd/gargantua/service"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/starter/starter_context"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"os"
	"sync"
	"text/tabwriter"
)

// usage:
// gargantua --config /path/to/config.yaml
func main() {
	// Define flags.
	fs := flag.NewFlagSet(common.ProjectName, flag.ExitOnError)
	var (
		configFile       = fs.String("config", "configs/config.yaml", "config path for gargantua")
		isNeedRestServer = fs.Bool("rest", true, "start a restful server while running grpc server")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("parse input params failed! error: %v", err)
	}

	// Load config.
	err = service.Init(*configFile)
	if err != nil {
		fmt.Printf("init gargantua grpc server failed! error: %v", err)
	}

	// get logger
	logger := common.GlobalLogger.Named("init")
	logger.Info("init gargantua grpc server success!")

	// set a channel to receive exit signal
	exitChan := make(chan struct{})

	// start gargantua servers
	wg := sync.WaitGroup{}

	// start grpc server
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := service.GargantuaGrpc()
		if err != nil {
			logger.Errorf("start gargantua grpc server failed! error: %v", err)
		}
	}()

	// start rest server
	if *isNeedRestServer {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := service.GargantuaRest()
			if err != nil {
				logger.Errorf("start gargantua rest server failed! error: %v", err)
			}
		}()
	}

	// wait for exit signal
	go func() {
		<-exitChan
		starter_context.CancelAllServer()
	}()

	// wait for all goroutines to finish
	wg.Wait()
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
