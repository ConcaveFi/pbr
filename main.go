package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	race "github.com/rrobrms/ethclient-providers-block-race"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Error: both a number of blocks and some websockets urls must be provided as arguments")
	}

	blockCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Error: blockCount must be a valid integer")
	}

	wsProviders := os.Args[2:]

	best, err := race.New(wsProviders, blockCount)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nyou should use this provider: %s\n", best)
}
