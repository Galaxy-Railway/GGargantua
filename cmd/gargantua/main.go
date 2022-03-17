package main

import (
	"flag"
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/cmd/gargantua/service"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"os"
	"text/tabwriter"
)

func main() {
	// Define flags.
	fs := flag.NewFlagSet(common.ProjectName, flag.ExitOnError)
	var (
		configFile = fs.String("config", "./config.yaml", "config path for gargantua")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	err := fs.Parse(os.Args[1:])
	if err != nil{
		fmt.Printf("parse input params failed! error: %v", err)
	}
	service.Gargantua(*configFile)
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