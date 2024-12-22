package main

import (
	"flag"
	"fmt"

	"github.com/debug-ing/codegate/internal"
)

func main() {
	mode := flag.String("mode", "init", "enter mode")
	flag.Parse()
	fmt.Println(*mode)
	if *mode == "init" {
		internal.Initialize()
		return
	}
	if *mode == "run" {
		internal.Run()
		return
	}
}
