package main

import (
	"log"

	"github.com/wushengyouya/tour/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

}
