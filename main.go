package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://localhost:22000")
	//address_erc20 := "0x9b2233dcfa0f3bf82ec734af762477eb2eda7b12"
	if err != nil {
		log.Fatal("client err", err)
	}
	checkTx(client, common.HexToHash("0x460ad592631cb7f3a65976732fdfbddb1f65c27c9eb3c743348b49da6b872638"))
	//deployErc20Contract(client)
}
