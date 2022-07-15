package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	ethClient, err := ethclient.Dial("http://localhost:22000")
	if err != nil {
		log.Fatal("client err", err)
	}

}

