package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func deployErc20Contract(client *ethclient.Client) { //部署erc20智能合约
	privateKey, err := crypto.HexToECDSA("1b2a4642eedbe3b657bc136e9639d3fed4c5748fcdfa548ae16e6df71d3c650e") //导入账户7的私钥
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public() //解析账户7的公钥
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) //解析账户7的地址
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0x00)
	gasLimit := uint64(3000000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	var code = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061061e806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063075461721461005157806327e235e31461006f57806340c10f191461009f578063d0679d34146100bb575b600080fd5b6100596100d7565b604051610066919061042a565b60405180910390f35b6100896004803603810190610084919061039f565b6100fb565b604051610096919061047c565b60405180910390f35b6100b960048036038101906100b491906103cc565b610113565b005b6100d560048036038101906100d091906103cc565b6101c5565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461016b57600080fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546101ba91906104c0565b925050819055505050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111561028a5780600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040517fcf479181000000000000000000000000000000000000000000000000000000008152600401610281929190610497565b60405180910390fd5b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102d99190610516565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461032f91906104c0565b925050819055507f3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd334533838360405161036993929190610445565b60405180910390a15050565b600081359050610384816105ba565b92915050565b600081359050610399816105d1565b92915050565b6000602082840312156103b5576103b46105b5565b5b60006103c384828501610375565b91505092915050565b600080604083850312156103e3576103e26105b5565b5b60006103f185828601610375565b92505060206104028582860161038a565b9150509250929050565b6104158161054a565b82525050565b6104248161057c565b82525050565b600060208201905061043f600083018461040c565b92915050565b600060608201905061045a600083018661040c565b610467602083018561040c565b610474604083018461041b565b949350505050565b6000602082019050610491600083018461041b565b92915050565b60006040820190506104ac600083018561041b565b6104b9602083018461041b565b9392505050565b60006104cb8261057c565b91506104d68361057c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561050b5761050a610586565b5b828201905092915050565b60006105218261057c565b915061052c8361057c565b92508282101561053f5761053e610586565b5b828203905092915050565b60006105558261055c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6105c38161054a565b81146105ce57600080fd5b50565b6105da8161057c565b81146105e557600080fd5b5056fea2646970667358221220d5f75cb414a70687adf45ee53b53fe847972893e832f2348c6f6dc1710054d4664736f6c63430008070033"
	data := []byte(code)
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
