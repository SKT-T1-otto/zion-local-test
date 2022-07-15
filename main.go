package main

import (
	"awesomeProject/store"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	ethClient, err := ethclient.Dial("http://localhost:22000")
	if err != nil {
		log.Fatal("client err", err)
	}
	////readByteCode(ethClient)
	instance1 := loadStore(ethClient)
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], "111")
	copy(value[:], "111")
	//writeContract(instance1, ethClient, key, value)
	checkContract(instance1, key)
	//for i := 0; i < 10; i++ {
	//	sendTx(ethClient)
	//}
	//queryStore(instance1)
	//deployContract(ethClient)
	//checkTx(ethClient, common.HexToHash("0xd190908afe146705a822321253058fe166b85aea316876fefa68e6accf501c6d"))
}
func queryStore(instance *store.Store) {
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"

}
func readByteCode(client *ethclient.Client) {
	contractAddress := common.HexToAddress("0xa75D692a249Df5C6465C0BA5C33e10d0Ba7F55f3")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode))
}
func loadStore(client *ethclient.Client) *store.Store {
	address := common.HexToAddress("0xa75D692a249Df5C6465C0BA5C33e10d0Ba7F55f3")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	return instance
}
func checkTx(client *ethclient.Client, addr common.Hash) {
	res1, isPending, err := client.TransactionByHash(context.Background(), addr)
	str, _ := json.MarshalIndent(res1, "", "	")
	fmt.Println(string(str))
	fmt.Println(res1, isPending, err)
	res2, err := client.TransactionReceipt(context.Background(), addr)
	fmt.Println(res2, err)
}

func deployContract(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("1b2a4642eedbe3b657bc136e9639d3fed4c5748fcdfa548ae16e6df71d3c650e")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	var chainID *big.Int = big.NewInt(60801)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}

func sendTx(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("4b0c9b9d685db17ac9f295cb12f9d7d2369f5bf524b3ce52ce424031cafda1ae")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(9000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x008473C8CadA4b807C1A27967D6E54c7c4AeD2e9")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
func writeContract(instance *store.Store, client *ethclient.Client, key [32]byte, value [32]byte) {
	privateKey, err := crypto.HexToECDSA("4b0c9b9d685db17ac9f295cb12f9d7d2369f5bf524b3ce52ce424031cafda1ae")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var chainID *big.Int = big.NewInt(60801)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

}
func checkContract(instance *store.Store, key [32]byte) {

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:]))

}
