package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	var version,help bool
	flag.BoolVar(&version, "version", false, "Print version information.")
	flag.BoolVar(&version, "v", false, "Print version information.")
	flag.BoolVar(&help,"help",false,"Print help information.")
	flag.BoolVar(&help,"h",false,"Print help information.")
	flag.Parse()
	if len(os.Args) <= 1 {
		fmt.Println("Need at least 1 parameter, try -h or --help for more information.")
		return
	}
	if version {
		fmt.Printf("MIGA v1.0.0-%s-%s\n",runtime.GOOS,runtime.GOARCH)
		return
	}
	if help {
		flag.Usage()
	}
}