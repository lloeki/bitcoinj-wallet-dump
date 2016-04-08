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

	for i, key := range w.Key {
		fmt.Printf("%d Publ: %s\n", i, base58.Check(key.PublicKey, base58.BITCOIN_PUBKEY))
		fmt.Printf("%d Priv: %s\n", i, base58.Check(key.SecretBytes, base58.BITCOIN_PRIVKEY))
	}
}
