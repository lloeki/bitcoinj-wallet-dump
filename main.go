package main

import (
	"fmt"
	"io/ioutil"
	"local/bitcoinj-wallet-dump/base58"
	"local/bitcoinj-wallet-dump/wallet"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	w := &wallet.Wallet{}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("loading error: ", err)
	}
	err = proto.Unmarshal(data, w)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Printf("Public key (WIF): %s\n", base58.Check(w.Key[0].PublicKey, base58.BITCOIN_PUBKEY))
	fmt.Printf("Private key (WIF): %s\n", base58.Check(w.Key[0].SecretBytes, base58.BITCOIN_PRIVKEY))
}
