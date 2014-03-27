package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagSet = flag.NewFlagSet("", flag.ExitOnError)
	port    = flagSet.String("port", "8080", "Port to listen on")
)

func main() {
	flagSet.Parse(os.Args[1:])
	fmt.Println(*port)
}
